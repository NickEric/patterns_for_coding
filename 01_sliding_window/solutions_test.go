package sliding_window

import "testing"

// FindAveragesOfSubArrays
func FindAveragesOfSubArrays(K int, arr []int) []float64 {
	results := []float64{}
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			result := float64(windowsSum) / float64(K)
			results = append(results, result)
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return results
}

func TestFindAveragesOfSubArrays(t *testing.T) {
	// 时间复杂度为O(n), 因为每一个元素都只被加减1次
	results := FindAveragesOfSubArrays(5, []int{1, 3, 2, 6, -1, 4, 1, 8, 2})
	t.Log("Averages of subarrays of size K: ", results)
}

// FindMaxSubArrayOfSizeK
func FindMaxSubArrayOfSizeK(K int, arr []int) int {
	max := 0
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			if windowsSum > max {
				max = windowsSum
			}
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return max
}

func TestFindMaxSubArrayOfSizeK(t *testing.T) {
	t.Log(FindMaxSubArrayOfSizeK(2, []int{2, 3, 4, 1, 5})) // 预期结果为7
}

// SmallestSubarrayWithGivenSum
func SmallestSubarrayWithGivenSum(arr []int, s int) int {
	sumShortest := len(arr) + 1
	sumLen := 0
	windowsSum, windowsStart := 0, 0

	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		sumLen += 1
		for windowsSum >= s {
			if sumLen < sumShortest {
				sumShortest = sumLen
			}
			sumLen -= 1
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}

	if sumShortest == len(arr)+1 {
		return 0
	}

	return sumShortest
}

func TestSmallestSubarrayWithGivenSum(t *testing.T) {
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 3, 2}, 7)) // 预期：2
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 8}, 7))    // 预期：1
	t.Log(SmallestSubarrayWithGivenSum([]int{3, 4, 1, 1, 6}, 8))    // 预期：3
}

// LongestSubstringKDistinct
func LongestSubstringKDistinct(s string, k int) int {
	longest := 0
	windowsStart := 0
	windows := make(map[byte]int)
	for windowsEnd := range s {
		windows[s[windowsEnd]] += 1
		if len(windows) <= k && windowsEnd-windowsStart+1 > longest {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windows) > k && windowsStart <= windowsEnd {
			if v, ok := windows[s[windowsStart]]; ok && v == 1 {
				delete(windows, s[windowsStart])
			} else {
				windows[s[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestLongestSubstringKDistinct(t *testing.T) {
	t.Log(LongestSubstringKDistinct("araaci", 2)) // 预期：4
	t.Log(LongestSubstringKDistinct("raa", 1))    // 预期：2
	t.Log(LongestSubstringKDistinct("cbbebi", 3)) // 预期：5
}

// Fruits into Baskets
func FruitsIntoBaskets(fruits []byte) int {
	longest := 0
	windowsMap := make(map[byte]int)
	windowsStart := 0
	for windowsEnd, f := range fruits {
		windowsMap[f] += 1
		if len(windowsMap) <= 2 && longest < windowsEnd-windowsStart+1 {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windowsMap) > 2 {
			if v, ok := windowsMap[fruits[windowsStart]]; ok && v == 1 {
				delete(windowsMap, fruits[windowsStart])
			} else {
				windowsMap[fruits[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestFruitsIntoBaskets(t *testing.T) {
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'A', 'C'}))      // 预期：3
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'B', 'B', 'C'})) // 预期：5
}

// NoRepeatSubstring
func NoRepeatSubstring(s string) int {
	windowsStart := 0
	windowsMap := make(map[byte]int) // key是字符，int是字符在字符串中的位置
	maxLength := 0
	for windowsEnd := range s {
		rightChar := s[windowsEnd] // 也就是滑窗右边的值
		// 如果这个rightChar已经再windows中存在了，那么就执行shrink操作
		if index, ok := windowsMap[rightChar]; ok {
			// 此处的shrink也是一个tricky，和之前的问题不同，此处的windowsStart可以直接的shrink到字符串中，上一个rightChar出现的位置的下一个位置
			// 此处还有一个tricky，我们没有实际的在map讲shrink过的字符删除，判断一个字符是否被删除，只要判断它的index是否大于windowsStart即可
			if windowsStart <= index {
				windowsStart = index + 1
			}
		}
		windowsMap[rightChar] = windowsEnd
		if windowsEnd-windowsStart+1 > maxLength {
			maxLength = windowsEnd - windowsStart + 1
		}
	}
	return maxLength
}

func TestNoRepeatSubstring(t *testing.T) {
	t.Log(NoRepeatSubstring("aabccbb")) // 预期：3
	t.Log(NoRepeatSubstring("abbbb"))   // 预期：2
	t.Log(NoRepeatSubstring("abccde"))  // 预期：3
}

// CharacterReplacement
// leetcode的一道类似的题 https://leetcode-cn.com/problems/longest-repeating-character-replacement/solution/tong-guo-ci-ti-liao-jie-yi-xia-shi-yao-shi-hua-don/
// 本题难点在于： maxRepeatLetterNumber 这个变量，一开始容易理解为 当前窗口中的最大重复字符数； 但实际这个变量的含义为，所有满足条件的窗口中，最大重复字符数
// 在本题中，因为我们只对“最长，有效的子字符串”感兴趣，所以窗口其实并没有shrink（收缩），严格来讲除了expand（扩展），就是做了shift（平移，即整体向右边移动的一格）
// 而平移过程中，当前窗口可能会覆盖到“无效”的子字符串（即不满足 windowsLength - maxRepeatChar <= k）
// 按理来收，每次平移之后，需要重新计算当前windows中的maxRepeatedChar，但是maxrepeatChar准确来说是历史上最大的重复值
// shrink 仅会发生在 maxRepeatChar 没有更新（即变得更大的时候), 且当前窗口长度-maxRepeatChar > k 的时候
func CharacterReplacement(s string, k int) int {
	windowsStart := 0
	longest := 0
	maxRepeatLetterNumber := 0
	windowsMap := make(map[byte]int) // key: 字符串中的字符 value: 字符在window中出现的次数
	sbyte := []byte(s)
	for windowEnd, b := range sbyte {
		windowsMap[b] += 1
		if windowsMap[b] >= maxRepeatLetterNumber {
			maxRepeatLetterNumber = windowsMap[b]
		}

		// 否则则有机会执行expand
		if windowEnd-windowsStart+1-maxRepeatLetterNumber > k {
			windowsMap[sbyte[windowsStart]] -= 1
			windowsStart += 1
			continue
		}

		// 仅当expand的时候，longest可能会更新
		if longest < windowEnd-windowsStart+1 {
			longest = windowEnd - windowsStart + 1
		}
	}
	return longest
}

func TestCharacterReplacement(t *testing.T) {
	t.Log(CharacterReplacement("aabccbb", 2)) // 预期：5
	t.Log(CharacterReplacement("abbcb", 1))   // 预期：4
	t.Log(CharacterReplacement("abccde", 1))  // 预期：3
	t.Log(CharacterReplacement("baaab", 2))   // 预期：5
}
