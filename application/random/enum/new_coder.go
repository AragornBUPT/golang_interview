package enum

type NewCoderCategory int

const (
	AlgorithmLearning NewCoderCategory = iota
	AlgorithmInterview
)

func (ncc NewCoderCategory) String() string {
	return []string{"算法学习篇", "算法笔面试篇"}[ncc]
}

type NewCoderLearningType int

const (
	BeginnersGuide130 NewCoderLearningType = iota
	IntroductionToAlgorithms
	AdvancedAlgorithms
	MasteringAlgorithms
)

func (nclt NewCoderLearningType) String() string {
	return []string{"新手入门130", "算法入门", "算法进阶", "算法登峰"}[nclt]
}

type NewCoderInterviewType int

const (
	Top101InterviewQuestions NewCoderInterviewType = iota
	MustPracticeWrittenExamTemplates
	RealWrittenExamQuestionsFromTopTechCompanies
	InputOutputPractice
	HighFrequencyWrittenExamQuestions
)

func (ncit NewCoderInterviewType) String() string {
	return []string{"面试TOP101", "笔试模板必刷", "笔试大厂真题", "输入输出练习", "笔试高频题目"}[ncit]
}
