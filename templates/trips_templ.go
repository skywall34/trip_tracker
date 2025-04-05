// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"encoding/json"
	"fmt"
	"github.com/skywall34/trip-tracker/internal/api"
	"github.com/skywall34/trip-tracker/internal/models"
	"time"
)

func TripsPage() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex-1 w-full max-w-6xl mx-auto px-4 py-8\"><div class=\"text-center mb-8 mt-4\"><!-- Title and subtitle --><h1 class=\"text-5xl sm:text-4xl font-bold text-gray-800 mb-2\">Saving Your Trips Is Easier Now</h1><!-- Search Bar --><div class=\"flex justify-center mt-4\"><input type=\"text\" placeholder=\"Search for a trip\" class=\"w-2/3 sm:w-1/2 rounded-l-full px-4 py-2 border border-gray-300 focus:outline-none\"> <button class=\"rounded-r-full bg-green-500 text-white px-6 py-2 hover:bg-green-600 transition\">Search</button></div></div><!-- Tabs --><div class=\"flex border rounded-full overflow-hidden w-full max-w-sm mx-auto mb-6\"><button hx-get=\"/trips?past=false\" hx-target=\"#trips-list\" hx-trigger=\"load, click\" class=\"w-1/2 bg-white text-gray-700 py-2 font-semibold transition hover:bg-gray-100 focus:outline-none\">Upcoming</button> <button hx-get=\"/trips?past=true\" hx-target=\"#trips-list\" class=\"w-1/2 bg-gray-200 text-gray-600 py-2 font-semibold transition hover:bg-gray-300 focus:outline-none\">Past</button> <button id=\"add-trip-btn\" class=\"p-2 bg-gray-200 rounded-full hover:bg-gray-300 transition\"><img src=\"/static/images/add-circle-svgrepo-com.svg\" alt=\"Add\" class=\"w-6 h-6\"></button></div><!-- Hidden Create Trip Form --><div id=\"create-trip-form\" class=\"hidden fixed inset-0 flex items-center justify-center\"><div class=\"bg-white p-6 rounded-lg shadow-md w-96\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateTripForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div><!-- Trip Cards List TODO: Filter only by user and date range --><div id=\"trips-list\"><!-- This will be populated by HTMX, calling RenderTrips or RenderPastTrips --></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func RenderTrips(trips []models.Trip) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- Trip Filters TODO: Show Date Filter --><div class=\"text-center text-lg font-semibold mt-4 text-green-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(time.Now().Format("2 Jan 2006"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 63, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, " - ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(time.Now().AddDate(1, 0, 0).Format("2 Jan 2006"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 63, Col: 93}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div><div class=\"space-y-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, trip := range trips {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"relative\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs("trip-element-" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 69, Col: 75}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"><!-- Flight Card TODO: Reservation, Terminal, Gate --><div class=\"flex justify-between text-gray-500 text-sm my-2 font-bold\"><span></span> <span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(
				fmt.Sprintf(
					"%dh %dm",
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime)) * time.Second).Hours()),
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime))*time.Second).Minutes())%60))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 77, Col: 127}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</span></div><div class=\"bg-white text-gray-700 p-6 rounded-lg w-full text-center shadow-lg relative\"><!-- Icon in Top-Right --><button class=\"absolute top-3 right-3 text-gray-500 hover:text-gray-700\" hx-delete=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/trips?id=" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 83, Col: 138}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" hx-target=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs("#trip-element-" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 83, Col: 189}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" hx-swap=\"outerHTML\"><img src=\"/static/images/trash-svgrepo-com.svg\" alt=\"Delete\" class=\"w-6 h-6\"></button><div class=\"flex justify-between items-center\"><!-- Departure Time stored in UTC--><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Departure)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 89, Col: 97}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 90, Col: 159}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\">Loading...</span></div><span class=\"text-3xl font-bold text-green-600\">→</span><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Arrival)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 96, Col: 94}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.ArrivalTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 97, Col: 157}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\">Loading...</span></div></div><div class=\"grid grid-cols-4 text-sm text-gray-400 mt-4 text-center\"><span class=\"col-span-1\">Flight</span> <span class=\"col-span-1\">Reservation</span> <span class=\"col-span-1\">Terminal</span> <span class=\"col-span-1\">Gate</span></div><div class=\"grid grid-cols-4 text-sm font-semibold mt-1 text-center text-gray-700\"><span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(trip.FlightNumber)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 109, Col: 67}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Reservation)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 110, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Terminal)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 111, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Gate)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 112, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</span></div></div><!-- Check-In Section TODO: Status Backend --><div class=\"bg-[#DBF3F8] text-gray-700 p-6 rounded-lg w-full text-center shadow mt-1\"><div class=\"flex justify-between\"><div class=\"text-center w-full\"><span class=\"block text-sm font-bold\">Status</span> <span class=\"block text-sm font-bold\">On Time</span></div><div class=\"text-center w-full\"><span class=\"block text-sm font-bold\">Check-In At</span> <span class=\"block text-sm font-bold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime)-(24*60*60), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 125, Col: 172}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "\">Loading...</span></div></div><div class=\"w-full bg-gray-300 h-2 rounded-full overflow-hidden mt-2\"><div class=\"bg-green-500 h-full w-3/5\"></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if time.Now().Unix() > int64(trip.DepartureTime)-(24*60*60) && time.Now().Unix() < int64(trip.DepartureTime)-(90*60) {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<!-- Check-In Button Only If Within the time frame (24hrs before and up to 90 minutes before departure) --> <button class=\"w-full bg-green-500 text-white py-2 mt-2 rounded-lg font-semibold hover:bg-green-600 transition\">Check In</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func RenderPastTrips(trips []models.Trip) templ.Component {
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
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<div class=\"space-y-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, trip := range trips {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<div class=\"relative\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs("trip-element-" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 150, Col: 75}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "\"><div class=\"flex justify-between text-gray-500 text-sm my-2 font-bold\"><span></span> <span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(
				fmt.Sprintf(
					"%dh %dm",
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime)) * time.Second).Hours()),
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime))*time.Second).Minutes())%60))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 157, Col: 127}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "</span></div><div class=\"bg-white text-gray-700 p-6 rounded-lg w-full text-center shadow-lg relative\"><!-- Icon in Top-Right --><button class=\"absolute top-3 right-3 text-gray-500 hover:text-gray-700\" hx-delete=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var21 string
			templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs("/trips?id=" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 163, Col: 138}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "\" hx-target=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var22 string
			templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs("#trip-element-" + fmt.Sprint(trip.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 163, Col: 189}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "\" hx-swap=\"outerHTML\"><img src=\"/static/images/trash-svgrepo-com.svg\" alt=\"Delete\" class=\"w-6 h-6\"></button><div class=\"flex justify-between items-center\"><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Departure)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 168, Col: 97}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var24 string
			templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 169, Col: 159}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "\">Loading...</span></div><span class=\"text-3xl font-bold text-green-600\">→</span><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var25 string
			templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Arrival)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 175, Col: 94}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var26 string
			templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.ArrivalTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 176, Col: 157}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "\">Loading...</span></div></div><div class=\"grid grid-cols-4 text-sm text-gray-400 mt-4 text-center\"><span class=\"col-span-1\">Flight</span> <span class=\"col-span-1\">Reservation</span> <span class=\"col-span-1\">Terminal</span> <span class=\"col-span-1\">Gate</span></div><div class=\"grid grid-cols-4 text-sm font-semibold mt-1 text-center text-gray-700\"><span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var27 string
			templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(trip.FlightNumber)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 188, Col: 67}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var28 string
			templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Reservation)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 189, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var29 string
			templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Terminal)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 190, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var30 string
			templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Gate)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 191, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, "</span></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func CreateTripPage() templ.Component {
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
		templ_7745c5c3_Var31 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var31 == nil {
			templ_7745c5c3_Var31 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, "<div class=\"flex-1 w-full max-w-6xl mx-auto px-4 py-8\"><div class=\"text-center mb-8 mt-4\"><!-- Search Bar --><div class=\"flex justify-center mt-4\"><form hx-get=\"/api/flights\" hx-target=\"#results\" class=\"flex w-2/3 sm:w-1/2\"><input type=\"text\" name=\"flight_iata\" placeholder=\"Search for a flight\" class=\"flex-grow rounded-l-full px-4 py-2 border border-gray-300 focus:outline-none\"> <button type=\"submit\" class=\"rounded-r-full bg-green-500 text-white px-6 py-2 hover:bg-green-600 transition\">Search</button></form></div><!-- Somewhere on the page --><div id=\"results\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func TripForm(flights api.FlightsAPIResponse) templ.Component {
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
		templ_7745c5c3_Var32 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var32 == nil {
			templ_7745c5c3_Var32 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, flight := range flights.Data {

			// Helper to safely dereference *string
			str := func(ptr *string) string {
				if ptr != nil {
					return *ptr
				}
				return ""
			}

			vals := map[string]string{
				"departure":     flight.Departure.IATA,
				"arrival":       flight.Arrival.IATA,
				"departuretime": flight.Departure.Scheduled.Format("2006-01-02T15:04"),
				"arrivaltime":   flight.Arrival.Scheduled.Format("2006-01-02T15:04"),
				"airline":       flight.Airline.Name,
				"flightnumber":  flight.FlightInfo.Number,
				"reservation":   "",
				"terminal":      str(flight.Departure.Terminal),
				"gate":          str(flight.Departure.Gate),
				"timezone":      flight.Departure.Timezone,
			}

			hxValsJSONBytes, err := json.Marshal(vals)
			if err != nil {
				panic("failed to marshal hx-vals JSON: " + err.Error())
			}
			hxValsJSON := string(hxValsJSONBytes)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 39, "<div class=\"bg-white shadow-md rounded p-4 flex justify-between items-center\"><div><h3 class=\"text-xl font-semibold\">Flight ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var33 string
			templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs(flight.FlightInfo.IATA)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 263, Col: 80}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 40, "</h3><p class=\"text-gray-600\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var34 string
			templ_7745c5c3_Var34, templ_7745c5c3_Err = templ.JoinStringErrs(flight.Departure.Airport)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 264, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var34))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 41, " → ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var35 string
			templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(flight.Arrival.Airport)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 264, Col: 95}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 42, ", ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var36 string
			templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(flight.Departure.Scheduled.Format("Jan 2, 2006 3:04 PM"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 264, Col: 155}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 43, "</p></div><button hx-post=\"/trips\" hx-vals=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var37 string
			templ_7745c5c3_Var37, templ_7745c5c3_Err = templ.JoinStringErrs(hxValsJSON)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 268, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var37))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 44, "\" hx-swap=\"none\" class=\"bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition\">Add to My Flights</button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return nil
	})
}

func CreateTripForm() templ.Component {
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
		templ_7745c5c3_Var38 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var38 == nil {
			templ_7745c5c3_Var38 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 45, "<div id=\"create-trip\" class=\"bg-white p-6 rounded-lg shadow-md\"><h2 class=\"text-xl font-bold text-gray-800 mb-4\">Add a New Trip</h2><form hx-post=\"/trips\" hx-target=\"#trips-list\" hx-swap=\"beforebegin\"><div class=\"grid grid-cols-2 gap-4\"><div><label class=\"block text-sm font-semibold text-gray-600\">Departure</label> <input type=\"text\" name=\"departure\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter departure location\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Arrival</label> <input type=\"text\" name=\"arrival\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter arrival location\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Departure Time</label> <input type=\"datetime-local\" name=\"departuretime\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Arrival Time</label> <input type=\"datetime-local\" name=\"arrivaltime\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Airline</label> <input type=\"text\" name=\"airline\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter airline name\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Flight Number</label> <input type=\"text\" name=\"flightnumber\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter flight number\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Reservation</label> <input type=\"text\" name=\"reservation\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter reservation code\"></div><div><label class=\"block text-sm font-semibold text-gray-600\">Terminal</label> <input type=\"text\" name=\"terminal\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter terminal\"></div><div><label class=\"block text-sm font-semibold text-gray-600\">Gate</label> <input type=\"text\" name=\"gate\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter gate\"></div></div><input type=\"hidden\" name=\"timezone\" id=\"timezone\"><!-- Hidden field for timezone --><button type=\"submit\" class=\"mt-4 w-full bg-green-500 text-white py-2 rounded-lg font-semibold hover:bg-green-600 transition\">Submit Trip</button> <button type=\"button\" id=\"close-trip-form\" class=\"mt-4 w-full bg-gray-300 text-white py-2 rounded-lg font-semibold hover:bg-gray-600 transition\">Cancel</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
