// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col items-center text-center py-16\"><h1 class=\"text-5xl font-bold mt-6 leading-tight text-green-700\">Exploring a New City is Easier Now</h1><p class=\"text-lg font-bold text-gray-600 mt-4 mb-6\">Plan your trips effortlessly with Mia's Trips</p><button class=\"rounded-full bg-green-500 text-white px-6 py-2 hover:bg-green-600 transition\">Start Planning Now</button></div><!-- Single Card Example --><section class=\"max-w-7xl mx-auto px-4 py-8\"><!-- Tailwind grid classes to create columns and spacing--><div class=\"grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6\"><a href=\"/createtrip\" class=\"bg-white rounded-lg shadow-lg overflow-hidden transform rotate-0 hover:rotate-3 transition\"><img src=\"/static/images/home_pic1.jpg\" alt=\"Card 1\" class=\"w-full h-48 object-cover\"><div class=\"p-4\"><h2 class=\"text-xl font-semibold text-green-700\">Create Your First Trip</h2><p class=\"text-gray-600 mt-2\">Start Adding Your Trips and Destinations!</p></div></a><!-- Card 2 --><a href=\"/statistics\" class=\"bg-white rounded-lg shadow-lg overflow-hidden transform rotate-0 hover:rotate-3 transition\"><img src=\"/static/images/home_pic2.jpg\" alt=\"Card 2\" class=\"w-full h-48 object-cover\"><div class=\"p-4\"><h2 class=\"text-xl font-semibold text-green-700\">Track Your Trip Data</h2><p class=\"text-gray-600 mt-2\">See your progress as you go on trips and explore new places.</p></div></a><!-- Card 3 --><a href=\"/trips\" class=\"bg-white rounded-lg shadow-lg overflow-hidden transform rotate-0 hover:rotate-3 transition\"><img src=\"/static/images/home_pic3.jpg\" alt=\"Card 3\" class=\"w-full h-48 object-cover\"><div class=\"p-4\"><h2 class=\"text-xl font-semibold text-green-700\">Trip Tracker</h2><p class=\"text-gray-600 mt-2\">Keep track of your trips and get updated on your travel plans.</p></div></a><!-- Card 4 --><a href=\"/worldmap\" class=\"bg-white rounded-lg shadow-lg overflow-hidden transform rotate-0 hover:rotate-3 transition\"><img src=\"/static/images/home_pic4.jpg\" alt=\"Card 3\" class=\"w-full h-48 object-cover\"><div class=\"p-4\"><h2 class=\"text-xl font-semibold text-green-700\">See the World</h2><p class=\"text-gray-600 mt-2\">Explore the world map and visualize your trips on a global scale.</p></div></a></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
