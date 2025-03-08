// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func TaskForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form hx-post=\"/todos\" hx-target=\"#tasks\" hx-swap=\"beforeend\" hx-on::after-request=\"if(event.detail.successful) this.reset()\"><div class=\"space-y-2\"><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-3xl font-semibold\">Add Task</h2><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-3\"><label for=\"title\" class=\"block text-sm/6 font-medium\">Task</label><div class=\"mt-2\"><input class=\"input block w-full rounded-md px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\" type=\"text\" name=\"title\" placeholder=\"Task\"></div></div><div class=\"sm:col-span-3\"><label for=\"status\" class=\"block text-sm/6 font-medium\">Status</label><div class=\"mt-2\"><input class=\"input block w-full rounded-md px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\" type=\"text\" name=\"status\" placeholder=\"Status\"></div></div></div></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><button class=\"btn btn-primary rounded-md px-3 py-2 text-sm font-semibold shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\" type=\"submit\">Save</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
