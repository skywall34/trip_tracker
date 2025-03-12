// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex-1 w-full max-w-6xl mx-auto px-4 py-8\"><div class=\"text-center mb-8 mt-4\"><!-- Title and subtitle --><h1 class=\"text-5xl sm:text-4xl font-bold text-gray-800 mb-2\">Saving Your Trips Is Easier Now</h1><!-- Create Trip Form -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateTripForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Search Bar --><div class=\"flex justify-center mt-4\"><input type=\"text\" placeholder=\"Search for a trip\" class=\"w-2/3 sm:w-1/2 rounded-l-full px-4 py-2 border border-gray-300 focus:outline-none\"> <button class=\"rounded-r-full bg-green-500 text-white px-6 py-2 hover:bg-green-600 transition\">Search</button></div></div><!-- Tabs --><div class=\"flex border rounded-full overflow-hidden w-full max-w-sm mx-auto mb-6\"><button hx-get=\"/trips?&amp;past=false\" hx-target=\"#trips-list\" hx-trigger=\"load, click\" class=\"w-1/2 bg-white text-gray-700 py-2 font-semibold transition hover:bg-gray-100 focus:outline-none\">Upcoming</button> <button hx-get=\"/trips?&amp;past=true\" hx-target=\"#trips-list\" class=\"w-1/2 bg-gray-200 text-gray-600 py-2 font-semibold transition hover:bg-gray-300 focus:outline-none\">Past</button></div><!-- Trip Cards List TODO: Filter only by user and date range --><div id=\"trips-list\"><!-- This will be populated by HTMX, calling RenderTrips or RenderPastTrips --></div></div>")
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 53, Col: 40}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 53, Col: 93}
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<!-- Flight Card TODO: Reservation, Terminal, Gate --> <div class=\"flex justify-between text-gray-500 text-sm my-2 font-bold\"><span></span> <span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(
				fmt.Sprintf(
					"%dh %dm",
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime)) * time.Second).Hours()),
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime))*time.Second).Minutes())%60))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 66, Col: 123}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span></div><div class=\"bg-white text-gray-700 p-6 rounded-lg w-full text-center shadow-lg\"><div class=\"flex justify-between items-cente\"><!-- Departure Time stored in UTC--><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Departure)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 74, Col: 93}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 75, Col: 155}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\">Loading...</span></div><span class=\"text-3xl font-bold text-green-600\">→</span><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Arrival)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 81, Col: 90}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.ArrivalTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 82, Col: 153}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\">Loading...</span></div></div><div class=\"grid grid-cols-4 text-sm text-gray-400 mt-4 text-center\"><span class=\"col-span-1\">Flight</span> <span class=\"col-span-1\">Reservation</span> <span class=\"col-span-1\">Terminal</span> <span class=\"col-span-1\">Gate</span></div><div class=\"grid grid-cols-4 text-sm font-semibold mt-1 text-center text-gray-700\"><span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(trip.FlightNumber)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 94, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Reservation)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 95, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Terminal)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 96, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Gate)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 97, Col: 55}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</span></div></div><!-- Check-In Section TODO: Status Backend --> <div class=\"bg-[#DBF3F8] text-gray-700 p-6 rounded-lg w-full text-center shadow mt-1\"><div class=\"flex justify-between\"><div class=\"text-center w-full\"><span class=\"block text-sm font-bold\">Status</span> <span class=\"block text-sm font-bold\">On Time</span></div><div class=\"text-center w-full\"><span class=\"block text-sm font-bold\">Check-In At</span> <span class=\"block text-sm font-bold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime)-(24*60*60), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 110, Col: 168}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\">Loading...</span></div></div><div class=\"w-full bg-gray-300 h-2 rounded-full overflow-hidden mt-2\"><div class=\"bg-green-500 h-full\" style=\"width: 60%;\"></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if time.Now().Unix() > int64(trip.DepartureTime)-(24*60*60) && time.Now().Unix() < int64(trip.DepartureTime)-(90*60) {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<!-- Check-In Button Only If Within the time frame (24hrs before and up to 90 minutes before departure) --> <button class=\"w-full bg-green-500 text-white py-2 mt-2 rounded-lg font-semibold hover:bg-green-600 transition\">Check In</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</div>")
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
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<div class=\"space-y-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, trip := range trips {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<div class=\"flex justify-between text-gray-500 text-sm my-2 font-bold\"><span></span> <span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(
				fmt.Sprintf(
					"%dh %dm",
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime)) * time.Second).Hours()),
					int((time.Duration(int64(trip.ArrivalTime)-int64(trip.DepartureTime))*time.Second).Minutes())%60))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 140, Col: 123}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</span></div><div class=\"bg-white text-gray-700 p-6 rounded-lg w-full text-center shadow-lg\"><div class=\"flex justify-between items-cente\"><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Departure)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 147, Col: 93}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.DepartureTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 148, Col: 155}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "\">Loading...</span></div><span class=\"text-3xl font-bold text-green-600\">→</span><div class=\"text-center w-full\"><span class=\"block text-lg font-bold text-green-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Arrival)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 154, Col: 90}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</span> <span class=\"block text-xl font-semibold time-convert\" data-utc=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(trip.ArrivalTime), 0).UTC().Format(time.RFC3339))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 155, Col: 153}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "\">Loading...</span></div></div><div class=\"grid grid-cols-4 text-sm text-gray-400 mt-4 text-center\"><span class=\"col-span-1\">Flight</span> <span class=\"col-span-1\">Reservation</span> <span class=\"col-span-1\">Terminal</span> <span class=\"col-span-1\">Gate</span></div><div class=\"grid grid-cols-4 text-sm font-semibold mt-1 text-center text-gray-700\"><span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var21 string
			templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(trip.FlightNumber)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 167, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var22 string
			templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Reservation)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 168, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Terminal)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 169, Col: 59}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</span> <span class=\"col-span-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var24 string
			templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(trip.Gate)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/trips.templ`, Line: 170, Col: 55}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "</span></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
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
		templ_7745c5c3_Var25 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var25 == nil {
			templ_7745c5c3_Var25 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<div class=\"bg-white p-6 rounded-lg shadow-md\"><h2 class=\"text-xl font-bold text-gray-800 mb-4\">Add a New Trip</h2><form hx-post=\"/trips\" hx-target=\"#trips-list\" hx-swap=\"beforebegin\"><div class=\"grid grid-cols-2 gap-4\"><div><label class=\"block text-sm font-semibold text-gray-600\">Departure</label> <input type=\"text\" name=\"departure\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter departure location\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Arrival</label> <input type=\"text\" name=\"arrival\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter arrival location\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Departure Time</label> <input type=\"datetime-local\" name=\"departuretime\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Arrival Time</label> <input type=\"datetime-local\" name=\"arrivaltime\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Airline</label> <input type=\"text\" name=\"airline\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter airline name\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Flight Number</label> <input type=\"text\" name=\"flightnumber\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter flight number\" required></div><div><label class=\"block text-sm font-semibold text-gray-600\">Reservation</label> <input type=\"text\" name=\"reservation\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter reservation code\"></div><div><label class=\"block text-sm font-semibold text-gray-600\">Terminal</label> <input type=\"text\" name=\"terminal\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter terminal\"></div><div><label class=\"block text-sm font-semibold text-gray-600\">Gate</label> <input type=\"text\" name=\"gate\" class=\"w-full border rounded-lg px-4 py-2 focus:outline-none\" placeholder=\"Enter gate\"></div></div><input type=\"hidden\" name=\"timezone\" id=\"timezone\"><!-- Hidden field for timezone --><button type=\"submit\" class=\"mt-4 w-full bg-green-500 text-white py-2 rounded-lg font-semibold hover:bg-green-600 transition\">Submit Trip</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
