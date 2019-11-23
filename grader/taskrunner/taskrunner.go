package taskrunner

import (
	"github.com/sandykarunia/fudge/language"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io/ioutil"
	"strings"
)

// TaskRunner contains methods to run single task in grading process
//go:generate mockery -name=TaskRunner
type TaskRunner interface {
	// PrepareSandbox prepares the sandbox before use, returns the sandbox instance
	PrepareSandbox() (sandbox.Sandbox, error)

	// PrepareSubmissionCode writes submission code into the sandbox, returns created filename
	PrepareSubmissionCode(sb sandbox.Sandbox, submissionCode string) (filename string, err error)

	// CompileCode compiles the code in a file inside the sandbox with specified language
	// returns the compiled code filename
	CompileCode(sb sandbox.Sandbox, filename string, lang language.Language) (string, error)

	// FetchAndWriteToFile fetches contents from URL defined in urls, and write it to the files sequentially
	// returns the filename for each url, sequentially.
	FetchAndWriteToFile(sb sandbox.Sandbox, urls []string) ([]string, error)
}

type taskRunnerImpl struct {
	sbFactory   sandbox.Factory
	logger      logger.Logger
	utilsString utils.String
	sdkHTTP     sdk.HTTPFunctions
}

func (t *taskRunnerImpl) FetchAndWriteToFile(sb sandbox.Sandbox, urls []string) ([]string, error) {
	var outputFilename []string
	// run for each URL
	for _, url := range urls {
		// use anonymous function so we can use defer
		// takes in url as parameter
		// returns created filename and error
		fname, err := func(u string) (string, error) {
			// get the data
			resp, err := t.sdkHTTP.Get(u)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			// generate filename
			generatedFilename := t.utilsString.GenerateRandomAlphanumeric(32)

			// write the body to the file
			if err = sb.WriteFile(generatedFilename, resp.Body); err != nil {
				return "", err
			}

			return generatedFilename, nil
		}(url)

		if err != nil {
			t.logger.Error("Failed to fetch and write to file, err = %s", err.Error())
			return nil, err
		}

		outputFilename = append(outputFilename, fname)
	}
	t.logger.Info("Fetch and write to file run successfully")
	return outputFilename, nil
}

func (t *taskRunnerImpl) CompileCode(sb sandbox.Sandbox, filename string, lang language.Language) (string, error) {
	outputFilename := t.utilsString.GenerateRandomAlphanumeric(16)
	compileCmd, args := lang.CompileCmd(filename, outputFilename)
	// defaults to 10s, 128MB memory, 20MB file size
	if err := sb.Run(
		10000, 128*1024, 20*1024,
		"", "", "",
		compileCmd, args...,
	); err != nil {
		t.logger.Error("Failed to compile code, err = %s", err.Error())
		return "", err
	}
	t.logger.Info("Code compiled successfully")
	return outputFilename, nil
}

func (t *taskRunnerImpl) PrepareSandbox() (sandbox.Sandbox, error) {
	sb, err := t.sbFactory.NewPreparedSandbox()
	if err != nil {
		t.logger.Error("Failed to prepare sandbox, err = %s", err.Error())
		return nil, err
	}
	t.logger.Info("Sandbox prepared with box-id = %d", sb.GetID())
	return sb, nil
}

func (t *taskRunnerImpl) PrepareSubmissionCode(sb sandbox.Sandbox, submissionCode string) (filename string, err error) {
	submissionCodeFilename := t.utilsString.GenerateRandomAlphanumeric(16)
	if err = sb.WriteFile(submissionCodeFilename, ioutil.NopCloser(strings.NewReader(submissionCode))); err != nil {
		t.logger.Error("Failed to prepare submission code, err = %s", err.Error())
	}
	t.logger.Info("Submission code prepared with filename = %s", submissionCodeFilename)
	return submissionCodeFilename, err
}
