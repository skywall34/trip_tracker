(function () {
  var f;
  var o = "hx-target-";
  function g(e, r) {
    return e.substring(0, r.length) === r;
  }
  function s(e, r) {
    if (!e || !r) return null;
    var t = r.toString();
    var s = [
      t,
      t.substring(0, 2) + "*",
      t.substring(0, 2) + "x",
      t.substring(0, 1) + "*",
      t.substring(0, 1) + "x",
      t.substring(0, 1) + "**",
      t.substring(0, 1) + "xx",
      "*",
      "x",
      "***",
      "xxx",
    ];
    if (g(t, "4") || g(t, "5")) {
      s.push("error");
    }
    for (var n = 0; n < s.length; n++) {
      var i = o + s[n];
      var a = f.getClosestAttributeValue(e, i);
      if (a) {
        if (a === "this") {
          return f.findThisElement(e, i);
        } else {
          return f.querySelectorExt(e, a);
        }
      }
    }
    return null;
  }
  function n(e) {
    if (e.detail.isError) {
      if (htmx.config.responseTargetUnsetsError) {
        e.detail.isError = false;
      }
    } else if (htmx.config.responseTargetSetsError) {
      e.detail.isError = true;
    }
  }
  htmx.defineExtension("response-targets", {
    init: function (e) {
      f = e;
      if (htmx.config.responseTargetUnsetsError === undefined) {
        htmx.config.responseTargetUnsetsError = true;
      }
      if (htmx.config.responseTargetSetsError === undefined) {
        htmx.config.responseTargetSetsError = false;
      }
      if (htmx.config.responseTargetPrefersExisting === undefined) {
        htmx.config.responseTargetPrefersExisting = false;
      }
      if (htmx.config.responseTargetPrefersRetargetHeader === undefined) {
        htmx.config.responseTargetPrefersRetargetHeader = true;
      }
    },
    onEvent: function (e, r) {
      if (
        e === "htmx:beforeSwap" &&
        r.detail.xhr &&
        r.detail.xhr.status !== 200
      ) {
        if (r.detail.target) {
          if (htmx.config.responseTargetPrefersExisting) {
            r.detail.shouldSwap = true;
            n(r);
            return true;
          }
          if (
            htmx.config.responseTargetPrefersRetargetHeader &&
            r.detail.xhr.getAllResponseHeaders().match(/HX-Retarget:/i)
          ) {
            r.detail.shouldSwap = true;
            n(r);
            return true;
          }
        }
        if (!r.detail.requestConfig) {
          return true;
        }
        var t = s(r.detail.requestConfig.elt, r.detail.xhr.status);
        if (t) {
          n(r);
          r.detail.shouldSwap = true;
          r.detail.target = t;
        }
        return true;
      }
    },
  });
})();
