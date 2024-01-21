package difficulty

import (
	"math"

	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func CalculateEyupStars(beatmap osu_parser.OsuFile) float64 {
	totalHitObjects := len(beatmap.HitObjects.List)
	noteDensity := float64(totalHitObjects) / float64(beatmap.DrainLength)

	difficulty := float64(0)

	if float64(beatmap.HitObjects.CountSlider)/float64(totalHitObjects) < 0.1 {
		difficulty = beatmap.Difficulty.HPDrainRate + beatmap.Difficulty.OverallDifficulty + beatmap.Difficulty.CircleSize
	} else {
		bpm60 := (60000.0 / beatmap.TimingPoints.TimingPoints[0].BeatLength) / 60
		difficulty = float64(beatmap.Difficulty.HPDrainRate+beatmap.Difficulty.OverallDifficulty+beatmap.Difficulty.CircleSize+math.Max(0, (min(4, bpm60*beatmap.Difficulty.SliderMultiplier-1.5)*2.5))) * 0.75
	}

	stars := float64(0)

	if difficulty > 21 {
		stars = (math.Min(difficulty, 30)/3*4 + min(20-0.032*math.Pow(noteDensity-5, 4), 20)) / 10
	} else if noteDensity >= 2.5 {
		stars = (math.Min(difficulty, 18)/18*10 + math.Min(40-40/math.Pow(5, 3.5)*math.Pow((min(noteDensity, 5)-5), 4), 40)) / 10
	} else if noteDensity < 1 {
		stars = (math.Min(difficulty, 18)/18*10)/10 + 0.25
	} else {
		stars = (math.Min(difficulty, 18)/18*10 + math.Min(25*(noteDensity-1), 40)) / 10
	}

	return stars
}
