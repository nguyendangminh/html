package html

import (
	"errors"
	"io"
)

var defaultEngine *Engine

func SetDefaultEngine(engine *Engine) error {
	if engine == nil {
		return errors.New("given engine is nil")
	}

	// be sure the templates are loaded
	if err := engine.Load(); err != nil {
		return err
	}
	defaultEngine = engine
	return nil
}

// Render will execute the template name along with the given values.
func Render(out io.Writer, template string, binding interface{}, layout ...string) error {
	if defaultEngine == nil {
		return errors.New("default engine is not set yet")
	}
	return defaultEngine.Render(out, template, binding, layout...)
}
