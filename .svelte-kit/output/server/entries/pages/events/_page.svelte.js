import { W as getContext, _ as store_get, Z as ensure_array_like, a1 as stringify, X as escape_html, $ as unsubscribe_stores, a0 as bind_props, T as pop, R as push, a2 as fallback, S as setContext, a3 as add_styles, a4 as store_set, a5 as store_mutate } from "../../../chunks/index.js";
import { w as writable, d as derived, r as readable, g as get } from "../../../chunks/index2.js";
import { a as attr, g as getAgenda, S as SvelteMap } from "../../../chunks/index-server.js";
import "clsx";
async function tick() {
}
const events$1 = writable([]);
const eventsError = writable("");
const DAY_IN_SECONDS = 86400;
function createDate(input = void 0) {
  if (input !== void 0) {
    return input instanceof Date ? _fromLocalDate(input) : _fromISOString(input);
  }
  return _fromLocalDate(/* @__PURE__ */ new Date());
}
function createDuration(input) {
  if (typeof input === "number") {
    input = { seconds: input };
  } else if (typeof input === "string") {
    let seconds = 0, exp = 2;
    for (let part of input.split(":", 3)) {
      seconds += parseInt(part, 10) * Math.pow(60, exp--);
    }
    input = { seconds };
  } else if (input instanceof Date) {
    input = { hours: input.getUTCHours(), minutes: input.getUTCMinutes(), seconds: input.getUTCSeconds() };
  }
  let weeks = input.weeks || input.week || 0;
  return {
    years: input.years || input.year || 0,
    months: input.months || input.month || 0,
    days: weeks * 7 + (input.days || input.day || 0),
    seconds: (input.hours || input.hour || 0) * 60 * 60 + (input.minutes || input.minute || 0) * 60 + (input.seconds || input.second || 0),
    inWeeks: !!weeks
  };
}
function cloneDate(date) {
  return new Date(date.getTime());
}
function addDuration(date, duration, x = 1) {
  date.setUTCFullYear(date.getUTCFullYear() + x * duration.years);
  let month = date.getUTCMonth() + x * duration.months;
  date.setUTCMonth(month);
  month %= 12;
  if (month < 0) {
    month += 12;
  }
  while (date.getUTCMonth() !== month) {
    subtractDay(date);
  }
  date.setUTCDate(date.getUTCDate() + x * duration.days);
  date.setUTCSeconds(date.getUTCSeconds() + x * duration.seconds);
  return date;
}
function subtractDuration(date, duration, x = 1) {
  return addDuration(date, duration, -x);
}
function addDay(date, x = 1) {
  date.setUTCDate(date.getUTCDate() + x);
  return date;
}
function subtractDay(date, x = 1) {
  return addDay(date, -x);
}
function setMidnight(date) {
  date.setUTCHours(0, 0, 0, 0);
  return date;
}
function toLocalDate(date) {
  return new Date(
    date.getUTCFullYear(),
    date.getUTCMonth(),
    date.getUTCDate(),
    date.getUTCHours(),
    date.getUTCMinutes(),
    date.getUTCSeconds()
  );
}
function toISOString(date, len = 19) {
  return date.toISOString().substring(0, len);
}
function datesEqual(date1, ...dates2) {
  return dates2.every((date2) => date1.getTime() === date2.getTime());
}
function nextClosestDay(date, day) {
  let diff2 = day - date.getUTCDay();
  date.setUTCDate(date.getUTCDate() + (diff2 >= 0 ? diff2 : diff2 + 7));
  return date;
}
function prevClosestDay(date, day) {
  let diff2 = day - date.getUTCDay();
  date.setUTCDate(date.getUTCDate() + (diff2 <= 0 ? diff2 : diff2 - 7));
  return date;
}
function noTimePart(date) {
  return typeof date === "string" && date.length <= 10;
}
function copyTime(toDate, fromDate) {
  toDate.setUTCHours(fromDate.getUTCHours(), fromDate.getUTCMinutes(), fromDate.getUTCSeconds(), 0);
  return toDate;
}
function nextDate(date, duration) {
  addDuration(date, duration);
  return date;
}
function prevDate(date, duration, hiddenDays) {
  subtractDuration(date, duration);
  if (hiddenDays.length && hiddenDays.length < 7) {
    while (hiddenDays.includes(date.getUTCDay())) {
      subtractDay(date);
    }
  }
  return date;
}
function _fromLocalDate(date) {
  return new Date(Date.UTC(
    date.getFullYear(),
    date.getMonth(),
    date.getDate(),
    date.getHours(),
    date.getMinutes(),
    date.getSeconds()
  ));
}
function _fromISOString(str) {
  const parts = str.match(/\d+/g);
  return new Date(Date.UTC(
    Number(parts[0]),
    Number(parts[1]) - 1,
    Number(parts[2]),
    Number(parts[3] || 0),
    Number(parts[4] || 0),
    Number(parts[5] || 0)
  ));
}
function assign(...args) {
  return Object.assign(...args);
}
function keys(object) {
  return Object.keys(object);
}
function symbol() {
  return Symbol("ec");
}
function isArray(value) {
  return Array.isArray(value);
}
function isFunction(value) {
  return typeof value === "function";
}
const identity = (x) => x;
function debounce(fn, handle, queueStore) {
  queueStore.update((queue) => queue.set(handle, fn));
}
let payloadProp = symbol();
function setPayload(obj, payload) {
  obj[payloadProp] = payload;
}
function hasPayload(obj) {
  return !!obj?.[payloadProp];
}
function getPayload(obj) {
  return obj[payloadProp];
}
function createElement(tag, className, content, attrs = []) {
  let el = document.createElement(tag);
  el.className = className;
  if (typeof content == "string") {
    el.innerText = content;
  } else if (content.domNodes) {
    el.replaceChildren(...content.domNodes);
  } else if (content.html) {
    el.innerHTML = content.html;
  }
  for (let attr2 of attrs) {
    el.setAttribute(...attr2);
  }
  return el;
}
function getElementWithPayload(x, y, root = document) {
  for (let el of root.elementsFromPoint(x, y)) {
    if (hasPayload(el)) {
      return el;
    }
    if (el.shadowRoot) {
      let shadowEl = getElementWithPayload(x, y, el.shadowRoot);
      if (shadowEl) {
        return shadowEl;
      }
    }
  }
  return null;
}
function createView(view2, _viewTitle, _currentRange, _activeRange) {
  return {
    type: view2,
    title: _viewTitle,
    currentStart: _currentRange.start,
    currentEnd: _currentRange.end,
    activeStart: _activeRange.start,
    activeEnd: _activeRange.end,
    calendar: void 0
  };
}
function toViewWithLocalDates(view2) {
  view2 = assign({}, view2);
  view2.currentStart = toLocalDate(view2.currentStart);
  view2.currentEnd = toLocalDate(view2.currentEnd);
  view2.activeStart = toLocalDate(view2.activeStart);
  view2.activeEnd = toLocalDate(view2.activeEnd);
  return view2;
}
function listView(view2) {
  return view2.startsWith("list");
}
let eventId = 1;
function createEvents(input) {
  return input.map((event) => {
    let result = {
      id: "id" in event ? String(event.id) : `{generated-${eventId++}}`,
      resourceIds: toArrayProp(event, "resourceId").map(String),
      allDay: event.allDay ?? (noTimePart(event.start) && noTimePart(event.end)),
      start: createDate(event.start),
      end: createDate(event.end),
      title: event.title ?? "",
      editable: event.editable,
      startEditable: event.startEditable,
      durationEditable: event.durationEditable,
      display: event.display ?? "auto",
      extendedProps: event.extendedProps ?? {},
      backgroundColor: event.backgroundColor ?? event.color,
      textColor: event.textColor,
      classNames: toArrayProp(event, "className"),
      styles: toArrayProp(event, "style")
    };
    if (result.allDay) {
      setMidnight(result.start);
      let end = cloneDate(result.end);
      setMidnight(result.end);
      if (!datesEqual(result.end, end) || datesEqual(result.end, result.start)) {
        addDay(result.end);
      }
    }
    return result;
  });
}
function toArrayProp(input, propName) {
  let result = input[propName + "s"] ?? input[propName] ?? [];
  return isArray(result) ? result : [result];
}
function createEventSources(input) {
  return input.map((source) => ({
    events: source.events,
    url: source.url && source.url.trimEnd("&") || "",
    method: source.method && source.method.toUpperCase() || "GET",
    extraParams: source.extraParams || {}
  }));
}
function createEventChunk(event, start, end) {
  return {
    start: event.start > start ? event.start : start,
    end: event.end < end ? event.end : end,
    event
  };
}
function sortEventChunks(chunks) {
  chunks.sort((a, b) => a.start - b.start || b.event.allDay - a.event.allDay);
}
function createEventContent(chunk, displayEventEnd, eventContent, theme, _intlEventTime, _view) {
  let timeText = _intlEventTime.formatRange(
    chunk.start,
    displayEventEnd && chunk.event.display !== "pointer" && !chunk.zeroDuration ? copyTime(cloneDate(chunk.start), chunk.end) : chunk.start
  );
  let content;
  if (eventContent) {
    content = isFunction(eventContent) ? eventContent({
      event: toEventWithLocalDates(chunk.event),
      timeText,
      view: toViewWithLocalDates(_view)
    }) : eventContent;
  }
  if (content === void 0) {
    let domNodes;
    switch (chunk.event.display) {
      case "background":
        domNodes = [];
        break;
      case "pointer":
        domNodes = [createTimeElement(timeText, chunk, theme)];
        break;
      default:
        domNodes = [
          ...chunk.event.allDay ? [] : [createTimeElement(timeText, chunk, theme)],
          createElement("h4", theme.eventTitle, chunk.event.title)
        ];
    }
    content = { domNodes };
  }
  return [timeText, content];
}
function createTimeElement(timeText, chunk, theme) {
  return createElement(
    "time",
    theme.eventTime,
    timeText,
    [["datetime", toISOString(chunk.start)]]
  );
}
function createEventClasses(eventClassNames, event, _view) {
  let result = event.classNames;
  if (eventClassNames) {
    if (isFunction(eventClassNames)) {
      eventClassNames = eventClassNames({
        event: toEventWithLocalDates(event),
        view: toViewWithLocalDates(_view)
      });
    }
    result = [
      ...isArray(eventClassNames) ? eventClassNames : [eventClassNames],
      ...result
    ];
  }
  return result;
}
function toEventWithLocalDates(event) {
  return _cloneEvent(event, toLocalDate);
}
function _cloneEvent(event, dateFn) {
  event = assign({}, event);
  event.start = dateFn(event.start);
  event.end = dateFn(event.end);
  return event;
}
function prepareEventChunks(chunks, hiddenDays) {
  let longChunks = {};
  if (chunks.length) {
    sortEventChunks(chunks);
    let prevChunk;
    for (let chunk of chunks) {
      let dates = [];
      let date = setMidnight(cloneDate(chunk.start));
      while (chunk.end > date) {
        if (!hiddenDays.includes(date.getUTCDay())) {
          dates.push(cloneDate(date));
          if (dates.length > 1) {
            let key = date.getTime();
            if (longChunks[key]) {
              longChunks[key].chunks.push(chunk);
            } else {
              longChunks[key] = {
                sorted: false,
                chunks: [chunk]
              };
            }
          }
        }
        addDay(date);
      }
      if (dates.length) {
        chunk.date = dates[0];
        chunk.days = dates.length;
        chunk.dates = dates;
        if (chunk.start < dates[0]) {
          chunk.start = dates[0];
        }
        let maxEnd = addDay(cloneDate(dates.at(-1)));
        if (chunk.end > maxEnd) {
          chunk.end = maxEnd;
        }
      } else {
        chunk.date = setMidnight(cloneDate(chunk.start));
        chunk.days = 1;
        chunk.dates = [chunk.date];
      }
      if (prevChunk && datesEqual(prevChunk.date, chunk.date)) {
        chunk.prev = prevChunk;
      }
      prevChunk = chunk;
    }
  }
  return longChunks;
}
function runReposition(refs, data) {
  refs.length = data.length;
  let result = [];
  for (let ref of refs) {
    result.push(ref?.reposition?.());
  }
  return result;
}
function eventIntersects(event, start, end, resources) {
  if (event.start < end && event.end > start) {
    if (resources) {
      if (!isArray(resources)) {
        resources = [resources];
      }
      return resources.some((resource) => event.resourceIds.includes(resource.id));
    }
    return true;
  }
  return false;
}
function helperEvent(display) {
  return previewEvent(display) || ghostEvent(display) || pointerEvent(display);
}
function bgEvent(display) {
  return display === "background";
}
function previewEvent(display) {
  return display === "preview";
}
function ghostEvent(display) {
  return display === "ghost";
}
function pointerEvent(display) {
  return display === "pointer";
}
function btnTextMonth(text) {
  return btnText(text, "month");
}
function btnText(text, period) {
  return {
    ...text,
    next: "Next " + period,
    prev: "Previous " + period
  };
}
function themeView(view2) {
  return (theme) => ({ ...theme, view: view2 });
}
function createResources(input) {
  let result = [];
  _createResources(input, 0, result);
  return result;
}
function _createResources(input, level, flat) {
  let result = [];
  for (let item of input) {
    let resource = createResource(item);
    result.push(resource);
    flat.push(resource);
    let payload = {
      level,
      children: [],
      expanded: true,
      hidden: false
    };
    setPayload(resource, payload);
    if (item.children) {
      payload.children = _createResources(item.children, level + 1, flat);
    }
  }
  return result;
}
function createResource(input) {
  return {
    id: String(input.id),
    title: input.title || "",
    eventBackgroundColor: input.eventBackgroundColor,
    eventTextColor: input.eventTextColor,
    extendedProps: input.extendedProps ?? {}
  };
}
function resourceBackgroundColor(event, resources) {
  return findResource(event, resources)?.eventBackgroundColor;
}
function resourceTextColor(event, resources) {
  return findResource(event, resources)?.eventTextColor;
}
function findResource(event, resources) {
  return resources.find((resource) => event.resourceIds.includes(resource.id));
}
function intl(locale, format) {
  return derived([locale, format], ([$locale, $format]) => {
    let intl2 = isFunction($format) ? { format: $format } : new Intl.DateTimeFormat($locale, $format);
    return {
      format: (date) => intl2.format(toLocalDate(date))
    };
  });
}
function intlRange(locale, format) {
  return derived([locale, format], ([$locale, $format]) => {
    let formatRange;
    if (isFunction($format)) {
      formatRange = $format;
    } else {
      let intl2 = new Intl.DateTimeFormat($locale, $format);
      formatRange = (start, end) => {
        if (start <= end) {
          return intl2.formatRange(start, end);
        } else {
          let parts = intl2.formatRangeToParts(end, start);
          let result = "";
          let sources = ["startRange", "endRange"];
          let processed = [false, false];
          for (let part of parts) {
            let i = sources.indexOf(part.source);
            if (i >= 0) {
              if (!processed[i]) {
                result += _getParts(sources[1 - i], parts);
                processed[i] = true;
              }
            } else {
              result += part.value;
            }
          }
          return result;
        }
      };
    }
    return {
      formatRange: (start, end) => formatRange(toLocalDate(start), toLocalDate(end))
    };
  });
}
function _getParts(source, parts) {
  let result = "";
  for (let part of parts) {
    if (part.source == source) {
      result += part.value;
    }
  }
  return result;
}
function createOptions(plugins) {
  let options = {
    allDayContent: void 0,
    allDaySlot: true,
    buttonText: {
      today: "today"
    },
    customButtons: {},
    date: /* @__PURE__ */ new Date(),
    datesSet: void 0,
    dayHeaderFormat: {
      weekday: "short",
      month: "numeric",
      day: "numeric"
    },
    dayHeaderAriaLabelFormat: {
      dateStyle: "full"
    },
    displayEventEnd: true,
    duration: { weeks: 1 },
    events: [],
    eventAllUpdated: void 0,
    eventBackgroundColor: void 0,
    eventTextColor: void 0,
    eventClassNames: void 0,
    eventClick: void 0,
    eventColor: void 0,
    eventContent: void 0,
    eventDidMount: void 0,
    eventMouseEnter: void 0,
    eventMouseLeave: void 0,
    eventSources: [],
    eventTimeFormat: {
      hour: "numeric",
      minute: "2-digit"
    },
    filterEventsWithResources: false,
    filterResourcesWithEvents: false,
    firstDay: 0,
    flexibleSlotTimeLimits: false,
    // ec option
    headerToolbar: {
      start: "title",
      center: "",
      end: "today prev,next"
    },
    height: void 0,
    hiddenDays: [],
    highlightedDates: [],
    // ec option
    lazyFetching: true,
    loading: void 0,
    locale: void 0,
    nowIndicator: false,
    resourceLabelContent: void 0,
    resourceLabelDidMount: void 0,
    resources: [],
    selectable: false,
    scrollTime: "06:00:00",
    slotDuration: "00:30:00",
    slotEventOverlap: true,
    slotHeight: 24,
    // ec option
    slotLabelFormat: {
      hour: "numeric",
      minute: "2-digit"
    },
    slotMaxTime: "24:00:00",
    slotMinTime: "00:00:00",
    slotWidth: 72,
    theme: {
      allDay: "ec-all-day",
      active: "ec-active",
      bgEvent: "ec-bg-event",
      bgEvents: "ec-bg-events",
      body: "ec-body",
      button: "ec-button",
      buttonGroup: "ec-button-group",
      calendar: "ec",
      compact: "ec-compact",
      content: "ec-content",
      day: "ec-day",
      dayHead: "ec-day-head",
      days: "ec-days",
      event: "ec-event",
      eventBody: "ec-event-body",
      eventTime: "ec-event-time",
      eventTitle: "ec-event-title",
      events: "ec-events",
      extra: "ec-extra",
      handle: "ec-handle",
      header: "ec-header",
      hiddenScroll: "ec-hidden-scroll",
      highlight: "ec-highlight",
      icon: "ec-icon",
      line: "ec-line",
      lines: "ec-lines",
      nowIndicator: "ec-now-indicator",
      otherMonth: "ec-other-month",
      resource: "ec-resource",
      sidebar: "ec-sidebar",
      sidebarTitle: "ec-sidebar-title",
      today: "ec-today",
      time: "ec-time",
      title: "ec-title",
      toolbar: "ec-toolbar",
      view: "",
      weekdays: ["ec-sun", "ec-mon", "ec-tue", "ec-wed", "ec-thu", "ec-fri", "ec-sat"],
      withScroll: "ec-with-scroll"
    },
    titleFormat: {
      year: "numeric",
      month: "short",
      day: "numeric"
    },
    view: void 0,
    viewDidMount: void 0,
    views: {}
  };
  for (let plugin of plugins) {
    plugin.createOptions?.(options);
  }
  return options;
}
function createParsers(plugins) {
  let parsers = {
    date: (date) => setMidnight(createDate(date)),
    duration: createDuration,
    events: createEvents,
    eventSources: createEventSources,
    hiddenDays: (days2) => [...new Set(days2)],
    highlightedDates: (dates) => dates.map((date) => setMidnight(createDate(date))),
    resources: createResources,
    scrollTime: createDuration,
    slotDuration: createDuration,
    slotMaxTime: createDuration,
    slotMinTime: createDuration
  };
  for (let plugin of plugins) {
    plugin.createParsers?.(parsers);
  }
  return parsers;
}
function diff(options, prevOptions) {
  let diff2 = [];
  for (let key of keys(options)) {
    if (options[key] !== prevOptions[key]) {
      diff2.push([key, options[key]]);
    }
  }
  assign(prevOptions, options);
  return diff2;
}
function dayGrid(state) {
  return derived(state.view, ($view) => $view?.startsWith("dayGrid"));
}
function activeRange(state) {
  return derived(
    [state._currentRange, state.firstDay, state.slotMaxTime, state._dayGrid],
    ([$_currentRange, $firstDay, $slotMaxTime, $_dayGrid]) => {
      let start = cloneDate($_currentRange.start);
      let end = cloneDate($_currentRange.end);
      if ($_dayGrid) {
        prevClosestDay(start, $firstDay);
        nextClosestDay(end, $firstDay);
      } else if ($slotMaxTime.days || $slotMaxTime.seconds > DAY_IN_SECONDS) {
        addDuration(subtractDay(end), $slotMaxTime);
        let start2 = subtractDay(cloneDate(end));
        if (start2 < start) {
          start = start2;
        }
      }
      return { start, end };
    }
  );
}
function currentRange(state) {
  return derived(
    [state.date, state.duration, state.firstDay],
    ([$date, $duration, $firstDay]) => {
      let start = cloneDate($date), end;
      if ($duration.months) {
        start.setUTCDate(1);
      } else if ($duration.inWeeks) {
        prevClosestDay(start, $firstDay);
      }
      end = addDuration(cloneDate(start), $duration);
      return { start, end };
    }
  );
}
function viewDates(state) {
  return derived([state._activeRange, state.hiddenDays], ([$_activeRange, $hiddenDays]) => {
    let dates = [];
    let date = setMidnight(cloneDate($_activeRange.start));
    let end = setMidnight(cloneDate($_activeRange.end));
    while (date < end) {
      if (!$hiddenDays.includes(date.getUTCDay())) {
        dates.push(cloneDate(date));
      }
      addDay(date);
    }
    if (!dates.length && $hiddenDays.length && $hiddenDays.length < 7) {
      state.date.update((date2) => {
        while ($hiddenDays.includes(date2.getUTCDay())) {
          addDay(date2);
        }
        return date2;
      });
      dates = get(state._viewDates);
    }
    return dates;
  });
}
function viewTitle(state) {
  return derived(
    [state.date, state._activeRange, state._intlTitle, state._dayGrid],
    ([$date, $_activeRange, $_intlTitle, $_dayGrid]) => {
      return $_dayGrid ? $_intlTitle.formatRange($date, $date) : $_intlTitle.formatRange($_activeRange.start, subtractDay(cloneDate($_activeRange.end)));
    }
  );
}
function view(state) {
  return derived([state.view, state._viewTitle, state._currentRange, state._activeRange], (args) => createView(...args));
}
function events(state) {
  let _events = writable([]);
  let abortController;
  let fetching = 0;
  let debounceHandle = {};
  derived(
    [state.events, state.eventSources, state._activeRange, state._fetchedRange, state.lazyFetching, state.loading],
    (values, set) => debounce(() => {
      let [$events, $eventSources, $_activeRange, $_fetchedRange, $lazyFetching, $loading] = values;
      if (!$eventSources.length) {
        set($events);
        return;
      }
      if (!$_fetchedRange.start || $_fetchedRange.start > $_activeRange.start || $_fetchedRange.end < $_activeRange.end || !$lazyFetching) {
        if (abortController) {
          abortController.abort();
        }
        abortController = new AbortController();
        if (isFunction($loading) && !fetching) {
          $loading(true);
        }
        let stopLoading = () => {
          if (--fetching === 0 && isFunction($loading)) {
            $loading(false);
          }
        };
        let events2 = [];
        let failure = (e) => stopLoading();
        let success = (data) => {
          events2 = events2.concat(createEvents(data));
          set(events2);
          stopLoading();
        };
        let startStr = toISOString($_activeRange.start);
        let endStr = toISOString($_activeRange.end);
        for (let source of $eventSources) {
          if (isFunction(source.events)) {
            let result = source.events({
              start: toLocalDate($_activeRange.start),
              end: toLocalDate($_activeRange.end),
              startStr,
              endStr
            }, success, failure);
            if (result !== void 0) {
              Promise.resolve(result).then(success, failure);
            }
          } else {
            let params = isFunction(source.extraParams) ? source.extraParams() : assign({}, source.extraParams);
            params.start = startStr;
            params.end = endStr;
            params = new URLSearchParams(params);
            let url = source.url, headers = {}, body;
            if (["GET", "HEAD"].includes(source.method)) {
              url += (url.includes("?") ? "&" : "?") + params;
            } else {
              headers["content-type"] = "application/x-www-form-urlencoded;charset=UTF-8";
              body = String(params);
            }
            fetch(url, { method: source.method, headers, body, signal: abortController.signal, credentials: "same-origin" }).then((response) => response.json()).then(success).catch(failure);
          }
          ++fetching;
        }
        $_fetchedRange.start = $_activeRange.start;
        $_fetchedRange.end = $_activeRange.end;
      }
    }, debounceHandle, state._queue),
    []
  ).subscribe(_events.set);
  return _events;
}
function now() {
  return readable(createDate(), (set) => {
    let interval = setInterval(() => {
      set(createDate());
    }, 1e3);
    return () => clearInterval(interval);
  });
}
function today(state) {
  return derived(state._now, ($_now) => setMidnight(cloneDate($_now)));
}
class State {
  constructor(plugins, input) {
    plugins = plugins || [];
    let options = createOptions(plugins);
    let parsers = createParsers(plugins);
    options = parseOpts(options, parsers);
    input = parseOpts(input, parsers);
    for (let [option, value] of Object.entries(options)) {
      this[option] = writable(value);
    }
    this._queue = writable(/* @__PURE__ */ new Map());
    this._queue2 = writable(/* @__PURE__ */ new Map());
    this._tasks = /* @__PURE__ */ new Map();
    this._auxiliary = writable([]);
    this._dayGrid = dayGrid(this);
    this._currentRange = currentRange(this);
    this._activeRange = activeRange(this);
    this._fetchedRange = writable({ start: void 0, end: void 0 });
    this._events = events(this);
    this._now = now();
    this._today = today(this);
    this._intlEventTime = intlRange(this.locale, this.eventTimeFormat);
    this._intlSlotLabel = intl(this.locale, this.slotLabelFormat);
    this._intlDayHeader = intl(this.locale, this.dayHeaderFormat);
    this._intlDayHeaderAL = intl(this.locale, this.dayHeaderAriaLabelFormat);
    this._intlTitle = intlRange(this.locale, this.titleFormat);
    this._bodyEl = writable(void 0);
    this._scrollable = writable(false);
    this._viewTitle = viewTitle(this);
    this._viewDates = viewDates(this);
    this._view = view(this);
    this._viewComponent = writable(void 0);
    this._interaction = writable({});
    this._iEvents = writable([null, null]);
    this._iClasses = writable(identity);
    this._iClass = writable(void 0);
    this._set = (key, value) => {
      if (validKey(key, this)) {
        if (parsers[key]) {
          value = parsers[key](value);
        }
        this[key].set(value);
      }
    };
    this._get = (key) => validKey(key, this) ? get(this[key]) : void 0;
    for (let plugin of plugins) {
      plugin.createStores?.(this);
    }
    if (input.view) {
      this.view.set(input.view);
    }
    let views = /* @__PURE__ */ new Set([...keys(options.views), ...keys(input.views ?? {})]);
    for (let view2 of views) {
      let defOpts = mergeOpts(options, options.views[view2] ?? {});
      let opts = mergeOpts(defOpts, input, input.views?.[view2] ?? {});
      let component = opts.component;
      filterOpts(opts, this);
      for (let key of keys(opts)) {
        let { set, _set = set, ...rest } = this[key];
        this[key] = {
          // Set value in all views
          set: ["buttonText", "theme"].includes(key) ? (value) => {
            if (isFunction(value)) {
              let result = value(defOpts[key]);
              opts[key] = result;
              set(set === _set ? result : value);
            } else {
              opts[key] = value;
              set(value);
            }
          } : (value) => {
            opts[key] = value;
            set(value);
          },
          _set,
          ...rest
        };
      }
      this.view.subscribe((newView) => {
        if (newView === view2) {
          this._viewComponent.set(component);
          if (isFunction(opts.viewDidMount)) {
            tick().then(() => opts.viewDidMount(get(this._view)));
          }
          for (let key of keys(opts)) {
            this[key]._set(opts[key]);
          }
        }
      });
    }
  }
}
function parseOpts(opts, parsers) {
  let result = { ...opts };
  for (let key of keys(parsers)) {
    if (key in result) {
      result[key] = parsers[key](result[key]);
    }
  }
  if (opts.views) {
    result.views = {};
    for (let view2 of keys(opts.views)) {
      result.views[view2] = parseOpts(opts.views[view2], parsers);
    }
  }
  return result;
}
function mergeOpts(...args) {
  let result = {};
  for (let opts of args) {
    let override = {};
    for (let key of ["buttonText", "theme"]) {
      if (isFunction(opts[key])) {
        override[key] = opts[key](result[key]);
      }
    }
    result = {
      ...result,
      ...opts,
      ...override
    };
  }
  return result;
}
function filterOpts(opts, state) {
  keys(opts).filter((key) => !validKey(key, state) || key == "view").forEach((key) => delete opts[key]);
}
function validKey(key, state) {
  return state.hasOwnProperty(key) && key[0] !== "_";
}
function Buttons($$payload, $$props) {
  push();
  var $$store_subs;
  let buttons = $$props["buttons"];
  let {
    _currentRange,
    _viewTitle,
    buttonText,
    customButtons,
    date,
    duration,
    hiddenDays,
    theme,
    view: view2
  } = getContext("state");
  let today2 = setMidnight(createDate()), isToday;
  isToday = today2 >= store_get($$store_subs ??= {}, "$_currentRange", _currentRange).start && today2 < store_get($$store_subs ??= {}, "$_currentRange", _currentRange).end || null;
  const each_array = ensure_array_like(buttons);
  $$payload.out += `<!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let button = each_array[$$index];
    if (button == "title") {
      $$payload.out += "<!--[-->";
      $$payload.out += `<h2${attr("class", store_get($$store_subs ??= {}, "$theme", theme).title)}></h2>`;
    } else {
      $$payload.out += "<!--[!-->";
      if (button == "prev") {
        $$payload.out += "<!--[-->";
        $$payload.out += `<button${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).button)} ec-${stringify(button)}`)}${attr("aria-label", store_get($$store_subs ??= {}, "$buttonText", buttonText).prev)}${attr("title", store_get($$store_subs ??= {}, "$buttonText", buttonText).prev)}><i${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).icon)} ec-${stringify(button)}`)}></i></button>`;
      } else {
        $$payload.out += "<!--[!-->";
        if (button == "next") {
          $$payload.out += "<!--[-->";
          $$payload.out += `<button${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).button)} ec-${stringify(button)}`)}${attr("aria-label", store_get($$store_subs ??= {}, "$buttonText", buttonText).next)}${attr("title", store_get($$store_subs ??= {}, "$buttonText", buttonText).next)}><i${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).icon)} ec-${stringify(button)}`)}></i></button>`;
        } else {
          $$payload.out += "<!--[!-->";
          if (button == "today") {
            $$payload.out += "<!--[-->";
            $$payload.out += `<button${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).button)} ec-${stringify(button)}`)}${attr("disabled", isToday, true)}>${escape_html(store_get($$store_subs ??= {}, "$buttonText", buttonText)[button])}</button>`;
          } else {
            $$payload.out += "<!--[!-->";
            if (store_get($$store_subs ??= {}, "$customButtons", customButtons)[button]) {
              $$payload.out += "<!--[-->";
              $$payload.out += `<button${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).button)} ec-${stringify(button)}${stringify(store_get($$store_subs ??= {}, "$customButtons", customButtons)[button].active ? " " + store_get($$store_subs ??= {}, "$theme", theme).active : "")}`)}></button>`;
            } else {
              $$payload.out += "<!--[!-->";
              if (button != "") {
                $$payload.out += "<!--[-->";
                $$payload.out += `<button${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).button)}${stringify(store_get($$store_subs ??= {}, "$view", view2) === button ? " " + store_get($$store_subs ??= {}, "$theme", theme).active : "")} ec-${stringify(button)}`)}>${escape_html(store_get($$store_subs ??= {}, "$buttonText", buttonText)[button])}</button>`;
              } else {
                $$payload.out += "<!--[!-->";
              }
              $$payload.out += `<!--]-->`;
            }
            $$payload.out += `<!--]-->`;
          }
          $$payload.out += `<!--]-->`;
        }
        $$payload.out += `<!--]-->`;
      }
      $$payload.out += `<!--]-->`;
    }
    $$payload.out += `<!--]-->`;
  }
  $$payload.out += `<!--]-->`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  bind_props($$props, { buttons });
  pop();
}
function Toolbar($$payload, $$props) {
  push();
  var $$store_subs;
  let { headerToolbar, theme } = getContext("state");
  let sections = { start: [], center: [], end: [] };
  {
    for (let key of keys(sections)) {
      sections[key] = store_get($$store_subs ??= {}, "$headerToolbar", headerToolbar)[key].split(" ").map((group) => group.split(","));
    }
  }
  const each_array = ensure_array_like(keys(sections));
  $$payload.out += `<nav${attr("class", store_get($$store_subs ??= {}, "$theme", theme).toolbar)}><!--[-->`;
  for (let $$index_1 = 0, $$length = each_array.length; $$index_1 < $$length; $$index_1++) {
    let key = each_array[$$index_1];
    const each_array_1 = ensure_array_like(sections[key]);
    $$payload.out += `<div${attr("class", `ec-${stringify(key)}`)}><!--[-->`;
    for (let $$index = 0, $$length2 = each_array_1.length; $$index < $$length2; $$index++) {
      let buttons = each_array_1[$$index];
      if (buttons.length > 1) {
        $$payload.out += "<!--[-->";
        $$payload.out += `<div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).buttonGroup)}>`;
        Buttons($$payload, { buttons });
        $$payload.out += `<!----></div>`;
      } else {
        $$payload.out += "<!--[!-->";
        Buttons($$payload, { buttons });
      }
      $$payload.out += `<!--]-->`;
    }
    $$payload.out += `<!--]--></div>`;
  }
  $$payload.out += `<!--]--></nav>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Auxiliary($$payload, $$props) {
  push();
  var $$store_subs;
  let {
    datesSet,
    _auxiliary,
    _activeRange,
    _queue,
    _view
  } = getContext("state");
  let debounceHandle = {};
  function runDatesSet(_activeRange2) {
    if (isFunction(store_get($$store_subs ??= {}, "$datesSet", datesSet))) {
      debounce(
        () => store_get($$store_subs ??= {}, "$datesSet", datesSet)({
          start: toLocalDate(_activeRange2.start),
          end: toLocalDate(_activeRange2.end),
          startStr: toISOString(_activeRange2.start),
          endStr: toISOString(_activeRange2.end),
          view: toViewWithLocalDates(store_get($$store_subs ??= {}, "$_view", _view))
        }),
        debounceHandle,
        _queue
      );
    }
  }
  runDatesSet(store_get($$store_subs ??= {}, "$_activeRange", _activeRange));
  const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$_auxiliary", _auxiliary));
  $$payload.out += `<!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let component = each_array[$$index];
    $$payload.out += `<!---->`;
    component?.($$payload, {});
    $$payload.out += `<!---->`;
  }
  $$payload.out += `<!--]-->`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Calendar($$payload, $$props) {
  push();
  var $$store_subs;
  let plugins = fallback($$props["plugins"], () => [], true);
  let options = fallback($$props["options"], () => ({}), true);
  let state = new State(plugins, options);
  setContext("state", state);
  let {
    _viewComponent,
    _bodyEl,
    _interaction,
    _iClass,
    _events,
    _queue,
    _queue2,
    _tasks,
    _scrollable,
    date,
    duration,
    hiddenDays,
    height,
    theme,
    view: view2
  } = state;
  let prevOptions = { ...options };
  function setOption(name, value) {
    state._set(name, value);
    return this;
  }
  function getOption(name) {
    let value = state._get(name);
    return value instanceof Date ? toLocalDate(value) : value;
  }
  function refetchEvents() {
    state._fetchedRange.set({ start: void 0, end: void 0 });
    return this;
  }
  function getEvents() {
    return store_get($$store_subs ??= {}, "$_events", _events).map(toEventWithLocalDates);
  }
  function getEventById(id) {
    for (let event of store_get($$store_subs ??= {}, "$_events", _events)) {
      if (event.id == id) {
        return toEventWithLocalDates(event);
      }
    }
    return null;
  }
  function addEvent(event) {
    event = createEvents([event])[0];
    store_get($$store_subs ??= {}, "$_events", _events).push(event);
    store_set(_events, store_get($$store_subs ??= {}, "$_events", _events));
    return event;
  }
  function updateEvent(event) {
    for (let e of store_get($$store_subs ??= {}, "$_events", _events)) {
      if (e.id == event.id) {
        event = createEvents([event])[0];
        assign(e, event);
        store_set(_events, store_get($$store_subs ??= {}, "$_events", _events));
        return event;
      }
    }
    return null;
  }
  function removeEventById(id) {
    let idx = store_get($$store_subs ??= {}, "$_events", _events).findIndex((event) => event.id == id);
    if (idx >= 0) {
      store_get($$store_subs ??= {}, "$_events", _events).splice(idx, 1);
      store_set(_events, store_get($$store_subs ??= {}, "$_events", _events));
    }
    return this;
  }
  function getView() {
    return toViewWithLocalDates(get(state._view));
  }
  function unselect() {
    store_get($$store_subs ??= {}, "$_interaction", _interaction).action?.unselect();
    return this;
  }
  function dateFromPoint(x, y) {
    let dayEl = getElementWithPayload(x, y);
    if (dayEl) {
      let info = getPayload(dayEl)(x, y);
      info.date = toLocalDate(info.date);
      return info;
    }
    return null;
  }
  function next() {
    store_set(date, nextDate(store_get($$store_subs ??= {}, "$date", date), store_get($$store_subs ??= {}, "$duration", duration)));
    return this;
  }
  function prev() {
    store_set(date, prevDate(store_get($$store_subs ??= {}, "$date", date), store_get($$store_subs ??= {}, "$duration", duration), store_get($$store_subs ??= {}, "$hiddenDays", hiddenDays)));
    return this;
  }
  for (let [name, value] of diff(options, prevOptions)) {
    setOption(name, value);
  }
  $$payload.out += `<div${add_styles({
    height: store_get($$store_subs ??= {}, "$height", height)
  })}${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).calendar)} ${stringify(store_get($$store_subs ??= {}, "$theme", theme).view)}${stringify(store_get($$store_subs ??= {}, "$_scrollable", _scrollable) ? " " + store_get($$store_subs ??= {}, "$theme", theme).withScroll : "")}${stringify(store_get($$store_subs ??= {}, "$_iClass", _iClass) ? " " + store_get($$store_subs ??= {}, "$theme", theme)[store_get($$store_subs ??= {}, "$_iClass", _iClass)] : "")}`)}${attr("role", listView(store_get($$store_subs ??= {}, "$view", view2)) ? "list" : "table")}>`;
  Toolbar($$payload);
  $$payload.out += `<!----> <!---->`;
  store_get($$store_subs ??= {}, "$_viewComponent", _viewComponent)?.($$payload, {});
  $$payload.out += `<!----></div> `;
  Auxiliary($$payload);
  $$payload.out += `<!---->`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  bind_props($$props, {
    plugins,
    options,
    setOption,
    getOption,
    refetchEvents,
    getEvents,
    getEventById,
    addEvent,
    updateEvent,
    removeEventById,
    getView,
    unselect,
    dateFromPoint,
    next,
    prev
  });
  pop();
}
function days(state) {
  return derived([state.date, state.firstDay, state.hiddenDays], ([$date, $firstDay, $hiddenDays]) => {
    let days2 = [];
    let day = cloneDate($date);
    let max = 7;
    while (day.getUTCDay() !== $firstDay && max) {
      subtractDay(day);
      --max;
    }
    for (let i = 0; i < 7; ++i) {
      if (!$hiddenDays.includes(day.getUTCDay())) {
        days2.push(cloneDate(day));
      }
      addDay(day);
    }
    return days2;
  });
}
function Header($$payload, $$props) {
  push();
  var $$store_subs;
  let {
    theme,
    _intlDayHeader,
    _intlDayHeaderAL,
    _days
  } = getContext("state");
  const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$_days", _days));
  $$payload.out += `<div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).header)}><div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).days)} role="row"><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let day = each_array[$$index];
    $$payload.out += `<div${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).day)} ${stringify(store_get($$store_subs ??= {}, "$theme", theme).weekdays?.[day.getUTCDay()])}`)} role="columnheader"><span${attr("aria-label", store_get($$store_subs ??= {}, "$_intlDayHeaderAL", _intlDayHeaderAL).format(day))}></span></div>`;
  }
  $$payload.out += `<!--]--></div> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).hiddenScroll)}></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Event($$payload, $$props) {
  push();
  var $$store_subs;
  let chunk = $$props["chunk"];
  let longChunks = fallback($$props["longChunks"], () => ({}), true);
  let inPopup = fallback($$props["inPopup"], false);
  let dates = fallback($$props["dates"], () => [], true);
  let {
    dayMaxEvents,
    displayEventEnd,
    eventAllUpdated,
    eventBackgroundColor,
    eventTextColor,
    eventClick,
    eventColor,
    eventContent,
    eventClassNames,
    eventDidMount,
    eventMouseEnter,
    eventMouseLeave,
    resources,
    theme,
    _view,
    _intlEventTime,
    _interaction,
    _iClasses,
    _hiddenEvents,
    _popupDate,
    _tasks
  } = getContext("state");
  let el;
  let event;
  let classes;
  let style;
  let margin = 1;
  let display;
  let onclick;
  function createHandler(fn, display2) {
    return !helperEvent(display2) && isFunction(fn) ? (jsEvent) => fn({
      event: toEventWithLocalDates(event),
      el,
      jsEvent,
      view: toViewWithLocalDates(store_get($$store_subs ??= {}, "$_view", _view))
    }) : void 0;
  }
  function reposition() {
    {
      return;
    }
  }
  event = chunk.event;
  {
    display = event.display;
    let bgColor = event.backgroundColor || resourceBackgroundColor(event, store_get($$store_subs ??= {}, "$resources", resources)) || store_get($$store_subs ??= {}, "$eventBackgroundColor", eventBackgroundColor) || store_get($$store_subs ??= {}, "$eventColor", eventColor);
    let txtColor = event.textColor || resourceTextColor(event, store_get($$store_subs ??= {}, "$resources", resources)) || store_get($$store_subs ??= {}, "$eventTextColor", eventTextColor);
    if (bgEvent(display)) {
      style = `width:calc(${chunk.days * 100}% + ${chunk.days - 1}px);`;
    } else {
      let marginTop = margin;
      if (event._margin) {
        let [_margin, _dates] = event._margin;
        if (chunk.date >= _dates[0] && chunk.date <= _dates.at(-1)) {
          marginTop = _margin;
        }
      }
      style = `width:calc(${chunk.days * 100}% + ${(chunk.days - 1) * 7}px);margin-top:${marginTop}px;`;
    }
    if (bgColor) {
      style += `background-color:${bgColor};`;
    }
    if (txtColor) {
      style += `color:${txtColor};`;
    }
    style += event.styles.join(";");
    classes = [
      bgEvent(display) ? store_get($$store_subs ??= {}, "$theme", theme).bgEvent : store_get($$store_subs ??= {}, "$theme", theme).event,
      ...store_get($$store_subs ??= {}, "$_iClasses", _iClasses)([], event),
      ...createEventClasses(store_get($$store_subs ??= {}, "$eventClassNames", eventClassNames), event, store_get($$store_subs ??= {}, "$_view", _view))
    ].join(" ");
  }
  createEventContent(chunk, store_get($$store_subs ??= {}, "$displayEventEnd", displayEventEnd), store_get($$store_subs ??= {}, "$eventContent", eventContent), store_get($$store_subs ??= {}, "$theme", theme), store_get($$store_subs ??= {}, "$_intlEventTime", _intlEventTime), store_get($$store_subs ??= {}, "$_view", _view));
  onclick = createHandler(store_get($$store_subs ??= {}, "$eventClick", eventClick), display);
  $$payload.out += `<article${attr("class", classes)}${attr("style", style)}${attr("role", onclick ? "button" : void 0)}${attr("tabindex", onclick ? 0 : void 0)}><!---->`;
  store_get($$store_subs ??= {}, "$_interaction", _interaction).resizer?.($$payload, { start: true, event });
  $$payload.out += `<!----> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).eventBody)}></div> <!---->`;
  store_get($$store_subs ??= {}, "$_interaction", _interaction).resizer?.($$payload, { event });
  $$payload.out += `<!----></article>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  bind_props($$props, {
    chunk,
    longChunks,
    inPopup,
    dates,
    reposition
  });
  pop();
}
function Popup($$payload, $$props) {
  push();
  var $$store_subs;
  let {
    buttonText,
    theme,
    _interaction,
    _intlDayPopover,
    _popupDate,
    _popupChunks
  } = getContext("state");
  let style = "";
  if (store_get($$store_subs ??= {}, "$_popupChunks", _popupChunks)) ;
  const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$_popupChunks", _popupChunks));
  $$payload.out += `<div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).popup)}${attr("style", style)}><div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).dayHead)}><time${attr("datetime", toISOString(store_get($$store_subs ??= {}, "$_popupDate", _popupDate), 10))}></time> <a role="button" tabindex="0"${attr("aria-label", store_get($$store_subs ??= {}, "$buttonText", buttonText).close)}>×</a></div> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).events)}><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let chunk = each_array[$$index];
    Event($$payload, { chunk, inPopup: true });
  }
  $$payload.out += `<!--]--></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Day($$payload, $$props) {
  push();
  var $$store_subs;
  let date = $$props["date"];
  let chunks = $$props["chunks"];
  let bgChunks = $$props["bgChunks"];
  let longChunks = $$props["longChunks"];
  let iChunks = fallback($$props["iChunks"], () => [], true);
  let dates = $$props["dates"];
  let {
    date: currentDate,
    dayMaxEvents,
    highlightedDates,
    moreLinkContent,
    theme,
    _hiddenEvents,
    _intlDayCell,
    _popupDate,
    _popupChunks,
    _today,
    _interaction,
    _queue
  } = getContext("state");
  let dayChunks, dayBgChunks;
  let isToday;
  let otherMonth;
  let highlight;
  let hiddenEvents = /* @__PURE__ */ new Set();
  let showPopup;
  let refs = [];
  function setPopupChunks() {
    let nextDay = addDay(cloneDate(date));
    let chunks2 = dayChunks.concat(longChunks[date.getTime()]?.chunks || []);
    store_set(_popupChunks, chunks2.map((chunk) => assign({}, chunk, createEventChunk(chunk.event, date, nextDay), { days: 1, dates: [date] })).sort((a, b) => a.top - b.top));
  }
  function reposition() {
    runReposition(refs, dayChunks);
  }
  {
    dayChunks = [];
    dayBgChunks = bgChunks.filter((bgChunk) => datesEqual(bgChunk.date, date));
    hiddenEvents.clear();
    hiddenEvents = hiddenEvents;
    for (let chunk of chunks) {
      if (datesEqual(chunk.date, date)) {
        dayChunks.push(chunk);
      }
    }
  }
  store_mutate($$store_subs ??= {}, "$_hiddenEvents", _hiddenEvents, store_get($$store_subs ??= {}, "$_hiddenEvents", _hiddenEvents)[date.getTime()] = hiddenEvents);
  isToday = datesEqual(date, store_get($$store_subs ??= {}, "$_today", _today));
  {
    otherMonth = date.getUTCMonth() !== store_get($$store_subs ??= {}, "$currentDate", currentDate).getUTCMonth();
    highlight = store_get($$store_subs ??= {}, "$highlightedDates", highlightedDates).some((d) => datesEqual(d, date));
  }
  if (store_get($$store_subs ??= {}, "$_hiddenEvents", _hiddenEvents) && hiddenEvents.size) {
    let text = "+" + hiddenEvents.size + " more";
    if (store_get($$store_subs ??= {}, "$moreLinkContent", moreLinkContent)) {
      isFunction(store_get($$store_subs ??= {}, "$moreLinkContent", moreLinkContent)) ? store_get($$store_subs ??= {}, "$moreLinkContent", moreLinkContent)({ num: hiddenEvents.size, text }) : store_get($$store_subs ??= {}, "$moreLinkContent", moreLinkContent);
    }
  }
  showPopup = store_get($$store_subs ??= {}, "$_popupDate", _popupDate) && datesEqual(date, store_get($$store_subs ??= {}, "$_popupDate", _popupDate));
  if (showPopup && longChunks && dayChunks) {
    tick().then(setPopupChunks);
  }
  const each_array = ensure_array_like(dayBgChunks);
  const each_array_1 = ensure_array_like(dayChunks);
  $$payload.out += `<div${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).day)} ${stringify(store_get($$store_subs ??= {}, "$theme", theme).weekdays?.[date.getUTCDay()])}${stringify(isToday ? " " + store_get($$store_subs ??= {}, "$theme", theme).today : "")}${stringify(otherMonth ? " " + store_get($$store_subs ??= {}, "$theme", theme).otherMonth : "")}${stringify(highlight ? " " + store_get($$store_subs ??= {}, "$theme", theme).highlight : "")}`)} role="cell"><time${attr("class", store_get($$store_subs ??= {}, "$theme", theme).dayHead)}${attr("datetime", toISOString(date, 10))}></time> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).bgEvents)}><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let chunk = each_array[$$index];
    Event($$payload, { chunk });
  }
  $$payload.out += `<!--]--></div> `;
  if (iChunks[2] && datesEqual(iChunks[2].date, date)) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).events)}>`;
    Event($$payload, { chunk: iChunks[2] });
    $$payload.out += `<!----></div>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> `;
  if (iChunks[0] && datesEqual(iChunks[0].date, date)) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<div${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).events)} ${stringify(store_get($$store_subs ??= {}, "$theme", theme).preview)}`)}>`;
    Event($$payload, { chunk: iChunks[0] });
    $$payload.out += `<!----></div>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).events)}><!--[-->`;
  for (let i = 0, $$length = each_array_1.length; i < $$length; i++) {
    let chunk = each_array_1[i];
    Event($$payload, { chunk, longChunks, dates });
  }
  $$payload.out += `<!--]--></div> `;
  if (showPopup) {
    $$payload.out += "<!--[-->";
    Popup($$payload);
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> <div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).dayFoot)}>`;
  if (hiddenEvents.size) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<a role="button" tabindex="0" aria-haspopup="true"></a>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  bind_props($$props, {
    date,
    chunks,
    bgChunks,
    longChunks,
    iChunks,
    dates,
    reposition
  });
  pop();
}
function Week($$payload, $$props) {
  push();
  var $$store_subs;
  let dates = $$props["dates"];
  let {
    _events,
    _iEvents,
    _queue2,
    _hiddenEvents,
    resources,
    filterEventsWithResources,
    hiddenDays,
    theme
  } = getContext("state");
  let chunks, bgChunks, longChunks, iChunks = [];
  let start;
  let end;
  let refs = [];
  let debounceHandle = {};
  function reposition() {
    debounce(() => runReposition(refs, dates), debounceHandle, _queue2);
  }
  {
    start = dates[0];
    end = addDay(cloneDate(dates.at(-1)));
  }
  {
    chunks = [];
    bgChunks = [];
    for (let event of store_get($$store_subs ??= {}, "$_events", _events)) {
      if (eventIntersects(event, start, end, store_get($$store_subs ??= {}, "$filterEventsWithResources", filterEventsWithResources) ? store_get($$store_subs ??= {}, "$resources", resources) : void 0)) {
        let chunk = createEventChunk(event, start, end);
        if (bgEvent(event.display)) {
          if (event.allDay) {
            bgChunks.push(chunk);
          }
        } else {
          chunks.push(chunk);
        }
      }
    }
    prepareEventChunks(bgChunks, store_get($$store_subs ??= {}, "$hiddenDays", hiddenDays));
    longChunks = prepareEventChunks(chunks, store_get($$store_subs ??= {}, "$hiddenDays", hiddenDays));
    reposition();
  }
  iChunks = store_get($$store_subs ??= {}, "$_iEvents", _iEvents).map((event) => {
    let chunk;
    if (event && eventIntersects(event, start, end)) {
      chunk = createEventChunk(event, start, end);
      prepareEventChunks([chunk], store_get($$store_subs ??= {}, "$hiddenDays", hiddenDays));
    } else {
      chunk = null;
    }
    return chunk;
  });
  if (store_get($$store_subs ??= {}, "$_hiddenEvents", _hiddenEvents)) {
    tick().then(reposition);
  }
  const each_array = ensure_array_like(dates);
  $$payload.out += `<div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).days)} role="row"><!--[-->`;
  for (let i = 0, $$length = each_array.length; i < $$length; i++) {
    let date = each_array[i];
    Day($$payload, {
      date,
      chunks,
      bgChunks,
      longChunks,
      iChunks,
      dates
    });
  }
  $$payload.out += `<!--]--></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  bind_props($$props, { dates });
  pop();
}
function Body($$payload, $$props) {
  push();
  var $$store_subs;
  let {
    _bodyEl,
    _viewDates,
    _hiddenEvents,
    dayMaxEvents,
    hiddenDays,
    theme
  } = getContext("state");
  let weeks;
  let days2;
  {
    weeks = [];
    days2 = 7 - store_get($$store_subs ??= {}, "$hiddenDays", hiddenDays).length;
    store_set(_hiddenEvents, {});
    store_get($$store_subs ??= {}, "$dayMaxEvents", dayMaxEvents);
    for (let i = 0; i < store_get($$store_subs ??= {}, "$_viewDates", _viewDates).length / days2; ++i) {
      let dates = [];
      for (let j = 0; j < days2; ++j) {
        dates.push(store_get($$store_subs ??= {}, "$_viewDates", _viewDates)[i * days2 + j]);
      }
      weeks.push(dates);
    }
  }
  const each_array = ensure_array_like(weeks);
  $$payload.out += `<div${attr("class", `${stringify(store_get($$store_subs ??= {}, "$theme", theme).body)}${stringify(store_get($$store_subs ??= {}, "$dayMaxEvents", dayMaxEvents) === true ? " " + store_get($$store_subs ??= {}, "$theme", theme).uniform : "")}`)}><div${attr("class", store_get($$store_subs ??= {}, "$theme", theme).content)}><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let dates = each_array[$$index];
    Week($$payload, { dates });
  }
  $$payload.out += `<!--]--></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function View($$payload) {
  Header($$payload);
  $$payload.out += `<!----> `;
  Body($$payload);
  $$payload.out += `<!---->`;
}
const DayGrid = {
  createOptions(options) {
    options.dayMaxEvents = false;
    options.dayCellFormat = { day: "numeric" };
    options.dayPopoverFormat = { month: "long", day: "numeric", year: "numeric" };
    options.moreLinkContent = void 0;
    options.buttonText.dayGridMonth = "month";
    options.buttonText.close = "Close";
    options.theme.uniform = "ec-uniform";
    options.theme.dayFoot = "ec-day-foot";
    options.theme.popup = "ec-popup";
    options.view = "dayGridMonth";
    options.views.dayGridMonth = {
      buttonText: btnTextMonth,
      component: View,
      dayHeaderFormat: { weekday: "short" },
      dayHeaderAriaLabelFormat: { weekday: "long" },
      displayEventEnd: false,
      duration: { months: 1 },
      theme: themeView("ec-day-grid ec-month-view"),
      titleFormat: { year: "numeric", month: "long" }
    };
  },
  createStores(state) {
    state._days = days(state);
    state._intlDayCell = intl(state.locale, state.dayCellFormat);
    state._intlDayPopover = intl(state.locale, state.dayPopoverFormat);
    state._hiddenEvents = writable({});
    state._popupDate = writable(null);
    state._popupChunks = writable([]);
  }
};
function _page($$payload, $$props) {
  push();
  var $$store_subs;
  const pickColorsFrom = [
    "#F6DC00",
    "#FFFFFF",
    "#AB69AD",
    "#FAC000",
    "#E7A2FD",
    "#DE9800"
  ];
  let agendaColors = new SvelteMap();
  let eventsAgendas = new SvelteMap();
  let plugins = [DayGrid];
  let options = {
    view: "dayGridMonth",
    events: [],
    editable: false,
    eventContent: (info) => {
      let titles = info.event.title.split("|");
      let desc = titles[1].replace(/^[\\n]+/g, "").replace(/\\n/g, "&#013;");
      return {
        html: `<div style="${getEventColors(info.event)}"><span class="fs-5 me-2 fw-bold">${info.event.start.getHours()}:${info.event.start.getMinutes()}</span><span class="fs-5" title="${desc}">${titles[0]} <br/> <i class="fs-6">${titles[2]}</i></span></div>`
      };
    },
    eventBackgroundColor: "#593196"
  };
  function getEventColors(event) {
    let colors = [];
    let color = "";
    event.resourceIds.forEach((r) => {
      if (!agendaColors.get(r)) {
        agendaColors.set(r, pickColorsFrom[agendaColors.size]);
      }
      colors.push(agendaColors.get(r));
    });
    if (colors.length > 1) {
      let percent = 100 / colors.length;
      color = "color: transparent;background-clip: text;background-image: linear-gradient(to right,";
      colors.forEach((c) => {
        color += `${c} ${percent}%,`;
      });
      color = color.substring(0, color.length - 1);
      color += ");";
    } else {
      color = `color: ${colors[0]};`;
    }
    return color;
  }
  events$1.subscribe((values) => {
    let list = [];
    values.forEach((value) => {
      let e = {
        id: value.id,
        //allDay: false,
        start: new Date(value.start),
        end: new Date(value.end),
        title: `${value.name}|${value.description}|${value.location}`,
        //todo see with desc
        //editable: false,
        //startEditable: false,
        //durationEditable: false,
        resourceIds: value.agendaIds
        //todo see ??
      };
      list.push(e);
      value.agendaIds.forEach((rId) => {
        if (!eventsAgendas.get(rId)) {
          getAgenda(rId).then((data) => {
            eventsAgendas.set(rId, data);
          }).catch(() => {
          });
        }
      });
    });
    options.events = list;
  });
  const each_array = ensure_array_like(agendaColors);
  $$payload.out += `<div class="d-flex flex-column"><h1 class="text-center color-yellow mb-lg-5">Events</h1> `;
  if (store_get($$store_subs ??= {}, "$eventsError", eventsError)) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<div class="text-center text-danger form-text fs-5">An error occurred : ${escape_html(store_get($$store_subs ??= {}, "$eventsError", eventsError))}</div>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> <div class="d-flex mb-4 justify-content-center"><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let [r, c] = each_array[$$index];
    $$payload.out += `<div class="d-flex me-3"><span class="p-3 me-2"${attr("style", `background-color: ${c};`)}></span> <span>${escape_html(eventsAgendas.get(r) ? eventsAgendas.get(r).name : r)}</span></div>`;
  }
  $$payload.out += `<!--]--></div> `;
  Calendar($$payload, { plugins, options });
  $$payload.out += `<!----></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
export {
  _page as default
};
