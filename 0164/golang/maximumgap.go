package maximumgap

import (
	// "fmt"
	"math"
)

type Bucket struct {
	Min int
	Max int
}

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[1] > nums[0] {
			return nums[1] - nums[0]
		}
		return nums[0] - nums[1]
	}

	var max, min, size int = nums[0], nums[0], len(nums) - 1

	for _, num := range nums {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	if max == min {
		return 0
	}

	interval := (max - min) / size
	remainder := (max - min) % size
	if remainder > 0 {
		interval++
	}

	if interval == 0 {
		return interval
	}

	buckets := make([]Bucket, size)
	for i := 0; i < size; i++ {
		buckets[i].Max = math.MinInt32
		buckets[i].Min = math.MaxInt32
	}

	for _, num := range nums {
		if num == min || num == max {
			continue
		}

		bucketIdx := (num - min) / interval
		if num < buckets[bucketIdx].Min {
			buckets[bucketIdx].Min = num
		}

		if num > buckets[bucketIdx].Max {
			buckets[bucketIdx].Max = num
		}
	}

	var maxGap, diff int
	var prevBucket *Bucket
	for i, bucket := range buckets {
		// empty bucket
		// fmt.Printf("maxGap: %d\n", maxGap)
		if bucket.Min > bucket.Max {
			continue
		}

		if prevBucket == nil {
			maxGap = bucket.Min - min
			prevBucket = &buckets[i]
			continue
		}

		diff = bucket.Min - prevBucket.Max
		if diff > maxGap {
			maxGap = diff
		}
		prevBucket = &buckets[i]
	}

	if prevBucket != nil {
		diff = max - prevBucket.Max
		if diff > maxGap {
			maxGap = diff
		}
	} else {
		maxGap = max - min
	}
	return maxGap

}
