# Iterator state and invariants

No new public entity or stored state. Existing iterator year/month/day defines the recurrence period; existing rule dtstart location gates optimization; effective until bounds instants. Existing remain buffer must drain before terminal state. A safe strictly-after period bound transitions the empty iterator to finished and sets r.len to total. Other transitions and per-candidate inclusive checks remain unchanged. Fresh iterators use setter-rebuilt rule values; active-iterator concurrent mutation remains outside this mission.
