package trace

import (
	"fmt"
	"io"
)

// Tracer 是一个接口用来描述一个可以通过代码来追踪事件的对象。
type Tracer interface {
	Trace(...interface{})
}

// 新创建的Tracer会将输出写到指定的io.Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// tracer是用来写入到io.Writer的Tracer对象。
type tracer struct {
	out io.Writer
}

// Trace 写参数到Tracers　io.Writer.
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
