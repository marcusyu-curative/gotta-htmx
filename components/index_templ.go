// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "examples/htmx-go/models"

func Index(todos []models.ToDo) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><link href=\"./static/css/tailwind.css\" rel=\"stylesheet\"><script src=\"https://unpkg.com/htmx.org@2.0.4\" integrity=\"sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+\" crossorigin=\"anonymous\"></script><script defer src=\"https://unpkg.com/alpinejs@3.14.8/dist/cdn.min.js\" integrity=\"sha384-X9kJyAubVxnP0hcA+AMMs21U445qsnqhnUF8EBlEpP3a42Kh/JwWjlv2ZcvGfphb\" crossorigin=\"anonymous\"></script><script>\n\t\t\t\tfunction updatePlaceholder() {\n\t\t\t\tconst tasks = document\n\t\t\t\t\t.getElementById(\"tasks\")\n\t\t\t\t\t.querySelectorAll(\"li.task\");\n\t\t\t\tconsole.log(tasks);\n\t\t\t\tconst placeholder = document.getElementById(\"no-tasks\");\n\t\t\t\tif (tasks.length > 0) {\n\t\t\t\t\tplaceholder.setAttribute(\"hidden\", true);\n\t\t\t\t} else {\n\t\t\t\t\tplaceholder.removeAttribute(\"hidden\");\n\t\t\t\t}\n\t\t\t\t}\n\t\t\t</script></head><body hx-on::after-request=\"console.log(&#39;body.after-request&#39;);\" hx-on::after-swap=\"console.log(&#39;body.after-swap&#39;);\"><div class=\"mx-auto max-w-7xl px-4 py-24 sm:px-6 sm:py-32 lg:px-8\"><div class=\"mx-auto max-w-2xl\"><h1 class=\"text-3xl font-bold\">Your Tasks</h1><div class=\"w-full max-w-lg\"><ul id=\"tasks\" class=\"list bg-base-100 rounded-box shadow-md\" hx-on::after-request=\"console.log(&#39;after-request&#39;); updatePlaceholder();\" hx-on::after-swap=\"console.log(&#39;after-swap&#39;); updatePlaceholder();\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, task := range todos {
			templ_7745c5c3_Err = Task(task).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<li class=\"list-row\" id=\"no-tasks\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(todos) > 0 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, " hidden")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, ">No tasks found</li></ul></div><form hx-post=\"/todos\" hx-target=\"#tasks\" hx-swap=\"beforeend\" hx-on::after-request=\"if(event.detail.successful) this.reset()\"><div class=\"space-y-2\"><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base/7 font-semibold\">Add Task</h2><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-3\"><label for=\"title\" class=\"block text-sm/6 font-medium\">Task</label><div class=\"mt-2\"><input class=\"input block w-full rounded-md px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\" type=\"text\" name=\"title\" placeholder=\"Task\"></div></div><div class=\"sm:col-span-3\"><label for=\"status\" class=\"block text-sm/6 font-medium\">Status</label><div class=\"mt-2\"><input class=\"input block w-full rounded-md px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\" type=\"text\" name=\"status\" placeholder=\"Status\"></div></div></div></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><button class=\"btn btn-primary rounded-md px-3 py-2 text-sm font-semibold shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\" type=\"submit\">Save</button></div></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
