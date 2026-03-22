package game

type Journey struct {
    Phase         int
    CurrentTick int
    TicksPerPhase int
    Daycount      int
    Phases        []string
}

func NewJourney() *Journey {
    return &Journey{
        Phase:         0,
        CurrentTick:   0,
        TicksPerPhase: 10,
        Daycount:      1,
        Phases:        []string{"🌅 Morning", "☀️ Midday", "🌇 Afternoon", "🌙 Evening", "💤 Rest"},
    }
}

func (j *Journey) Tick() {
	j.CurrentTick++ // increment the tick counter

	// check if we've hit the ticks per phase threshold
	if j.CurrentTick == j.TicksPerPhase{

		j.Phase++ // advance phase
		j.CurrentTick = 0 // reset tick counter

		// check if past last phase
		if j.Phase == len(j.Phases) { 
			j.Phase = 0 // reset phase to morning
			j.Daycount++
		}
	}
}
