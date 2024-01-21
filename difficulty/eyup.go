package difficulty

import (
	"math"

	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func CalculateEyupStars(beatmap osu_parser.OsuFile) float64 {
	totalHitObjects := len(beatmap.HitObjects.List)

	if len(beatmap.TimingPoints.TimingPoints) == 0 || totalHitObjects == 0 {
		return 0
	}

	noteDensity := float64(totalHitObjects) / float64(beatmap.DrainLength)

	difficulty := float64(0)

	diffSettingsSum := beatmap.Difficulty.HPDrainRate + beatmap.Difficulty.OverallDifficulty + beatmap.Difficulty.CircleSize

	if float64(beatmap.HitObjects.CountSlider)/float64(totalHitObjects) < 0.1 {
		difficulty = diffSettingsSum
	} else {
		beatLength := beatmap.TimingPoints.TimingPoints[0].BeatLength
		difficulty = (diffSettingsSum + math.Max(0, (math.Min(4, 1000/beatLength*beatmap.Difficulty.SliderMultiplier-1.5)*2.5))) * 0.75
	}

	stars := float64(0)

	if difficulty > 21 {
		stars = (math.Min(difficulty, 30)/3*4 + math.Min(20-0.032*math.Pow(noteDensity-5, 4), 20)) / 10
	} else if noteDensity >= 2.5 {
		stars = (math.Min(difficulty, 18)/18*10 + math.Min(40-40/math.Pow(5, 3.5)*math.Pow((math.Min(noteDensity, 5)-5), 4), 40)) / 10
	} else if noteDensity < 1 {
		stars = (math.Min(difficulty, 18)/18*10)/10 + 0.25
	} else {
		stars = (math.Min(difficulty, 18)/18*10 + math.Min(25*(noteDensity-1), 40)) / 10
	}

	return math.Min(5, stars)
}
