package random

import (
	"fmt"
	"golang_interview/application/random/enum"
	"math/rand"
)

// leetcode
func RandLeetcode() {
	question_number := rand.Intn(4131) + 1

	fmt.Println(question_number)
}

// newcoder
func RandNewCoder() {
	categories := []enum.NewCoderCategory{
		enum.AlgorithmLearning,
		enum.AlgorithmInterview,
	}
	categoryNum := rand.Intn(len(categories))
	category := categories[categoryNum]
	fmt.Println(category)

	var count int
	switch category {
	case enum.AlgorithmLearning:
		types := []enum.NewCoderLearningType{
			enum.BeginnersGuide130,
			enum.IntroductionToAlgorithms,
			enum.AdvancedAlgorithms,
			enum.MasteringAlgorithms,
		}
		typeNum := rand.Intn(len(types))
		typeMy := types[typeNum]
		fmt.Println(typeMy)

		switch typeMy {
		case enum.BeginnersGuide130:
			count = 130
		case enum.IntroductionToAlgorithms:
			count = 86
		case enum.AdvancedAlgorithms:
			count = 102
		case enum.MasteringAlgorithms:
			count = 154
		}
	case enum.AlgorithmInterview:
		types := []enum.NewCoderInterviewType{
			enum.Top101InterviewQuestions,
			enum.MustPracticeWrittenExamTemplates,
			enum.RealWrittenExamQuestionsFromTopTechCompanies,
			enum.InputOutputPractice,
			enum.HighFrequencyWrittenExamQuestions,
		}
		typeNum := rand.Intn(len(types))
		typeMy := types[typeNum]
		fmt.Println(typeMy)

		switch typeMy {
		case enum.Top101InterviewQuestions:
			count = 101
		case enum.MustPracticeWrittenExamTemplates:
			count = 147
		case enum.RealWrittenExamQuestionsFromTopTechCompanies:
			count = 3758
		case enum.InputOutputPractice:
			count = 18
		case enum.HighFrequencyWrittenExamQuestions:
			count = 415
		}
	}
	questionNum := rand.Intn(count) + 1
	fmt.Println(questionNum)
}
