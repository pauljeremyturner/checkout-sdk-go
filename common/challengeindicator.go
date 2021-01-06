package common

// ChallengeIndicator ...
type ChallengeIndicator string

const (
	// Default ...
	Default ChallengeIndicator = "no_preference"
	// NoChallengeRequested ...
	NoChallengeRequested ChallengeIndicator = "no_challenge_requested"
	// ChallengeRequestByMerchant ...
	ChallengeRequestByMerchant ChallengeIndicator = "challenge_request_by_merchant"
	// ChallengeRequestAsMandate ...
	ChallengeRequestAsMandate ChallengeIndicator = "challenge_request_as_mandate"
)

func (c ChallengeIndicator) String() string {
	return string(c)
}
