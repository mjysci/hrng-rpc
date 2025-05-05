package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/mjysci/hrng-rpc/internal/config"
	"github.com/mjysci/hrng-rpc/internal/utils"
)

type GenerateIntegersArgs struct {
	APIKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Min         int64  `json:"min"`
	Max         int64  `json:"max"`
	Replacement *bool  `json:"replacement"`
	Base        int    `json:"base,omitempty"`
}

type GenerateIntegerSequencesArgs struct {
	APIKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Length      any    `json:"length"`
	Min         any    `json:"min"`
	Max         any    `json:"max"`
	Replacement any    `json:"replacement,omitempty"`
	Base        any    `json:"base,omitempty"`
}

type GenerateDecimalFractionsArgs struct {
	APIKey        string `json:"apiKey"`
	N             int    `json:"n"`
	DecimalPlaces int    `json:"decimalPlaces"`
	Replacement   *bool  `json:"replacement"`
}

type GenerateGaussiansArgs struct {
	APIKey            string  `json:"apiKey"`
	N                 int     `json:"n"`
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standardDeviation"`
	SignificantDigits int     `json:"significantDigits"`
}

type GenerateStringsArgs struct {
	APIKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Length      int    `json:"length"`
	Characters  string `json:"characters"`
	Replacement *bool  `json:"replacement"`
}

type GenerateUUIDsArgs struct {
	APIKey string `json:"apiKey"`
	N      int    `json:"n"`
}

type GenerateBlobsArgs struct {
	APIKey string `json:"apiKey"`
	N      int    `json:"n"`
	Size   int    `json:"size"`
	Format string `json:"format,omitempty"`
}

type ResultData struct {
	Data           any    `json:"data"` // Can be [][]int64 or []float64
	CompletionTime string `json:"completionTime"`
}

type RandomResult struct {
	Random   ResultData `json:"random"`
	BitsUsed int        `json:"bitsUsed"`
}

type RandomService struct {
	randomGenerator utils.RandomGenerator
	config          *config.Config
}

// helper to convert float64 to int64 safely
func sampleInt64(v float64) int64 {
	return int64(v + 0.5)
}

func NewRandomService(randomGenerator utils.RandomGenerator, config *config.Config) *RandomService {
	return &RandomService{
		randomGenerator: randomGenerator,
		config:          config,
	}
}

func (s *RandomService) validateAPIKey(key string) error {
	if len(s.config.Server.APIKey) == 0 {
		return nil // No API keys configured, accept all
	}
	for _, validKey := range s.config.Server.APIKey {
		if key == validKey {
			return nil
		}
	}
	return fmt.Errorf("invalid API key")
}

// generateRandomInt generates a random integer within [min, max] using the configured generator
func (s *RandomService) generateRandomInt(min, max int64) (int64, error) {
	return s.randomGenerator.GenerateInt(min, max)
}

// generates random integers within a user-defined range
func (s *RandomService) GenerateIntegers(r *http.Request, args *GenerateIntegersArgs, reply *RandomResult) error {
	if args.Replacement == nil {
		defaultReplacement := true
		args.Replacement = &defaultReplacement
	}
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 1e4 {
		return fmt.Errorf("parameter 'n' must be between 1 and 10000, got %d", args.N)
	}
	if args.Min < -1e9 || args.Min > 1e9 {
		return fmt.Errorf("parameter 'min' must be between -1e9 and 1e9, got %d", args.Min)
	}
	if args.Max < -1e9 || args.Max > 1e9 {
		return fmt.Errorf("parameter 'max' must be between -1e9 and 1e9, got %d", args.Max)
	}
	if args.Min > args.Max {
		return fmt.Errorf("parameter 'min' (%d) cannot be greater than 'max' (%d)", args.Min, args.Max)
	}
	if args.Base != 0 && args.Base != 2 && args.Base != 8 && args.Base != 10 && args.Base != 16 {
		return fmt.Errorf("parameter 'base' must be 2, 8, 10, or 16, got %d", args.Base)
	}

	rangeSize := args.Max - args.Min + 1
	if !*args.Replacement && int64(args.N) > rangeSize {
		return fmt.Errorf("cannot generate %d unique integers from a range of size %d (min: %d, max: %d)", args.N, rangeSize, args.Min, args.Max)
	}

	startTime := time.Now()
	data := make([]any, 0, args.N)
	bitsUsed := 0

	if *args.Replacement {
		for i := 0; i < args.N; i++ {
			num, err := s.generateRandomInt(args.Min, args.Max)
			if err != nil {
				log.Printf("Error generating random int: %v", err)
				return fmt.Errorf("failed to generate random number from system source: %w", err)
			}
			if args.Base == 2 {
				data = append(data, fmt.Sprintf("%b", num))
			} else if args.Base == 8 {
				data = append(data, fmt.Sprintf("%o", num))
			} else if args.Base == 16 {
				data = append(data, fmt.Sprintf("%x", num))
			} else {
				data = append(data, num)
			}
			bitsUsed += 64
		}
	} else {
		samples, err := utils.SampleHybrid(int(rangeSize), args.N)
		if err != nil {
			log.Printf("Error sampling numbers: %v", err)
			return fmt.Errorf("failed to generate unique random numbers: %w", err)
		}
		for _, offset := range samples {
			num := args.Min + int64(offset)
			if args.Base == 2 {
				data = append(data, fmt.Sprintf("%b", num))
			} else if args.Base == 8 {
				data = append(data, fmt.Sprintf("%o", num))
			} else if args.Base == 16 {
				data = append(data, fmt.Sprintf("%x", num))
			} else {
				data = append(data, num)
			}
		}
		bitsUsed = args.N * 64 // approximate bits used
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates uniform or multiform sequences of random integers within user-defined ranges
func (s *RandomService) GenerateIntegerSequences(r *http.Request, args *GenerateIntegerSequencesArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 1000 {
		return fmt.Errorf("parameter 'n' must be between 1 and 1000, got %d", args.N)
	}

	// Normalize lengths
	lengths := make([]int, args.N)
	switch v := args.Length.(type) {
	case float64:
		length := int(v)
		if length < 1 || length > 10000 {
			return fmt.Errorf("parameter 'length' must be between 1 and 10000, got %d", length)
		}
		for i := range lengths {
			lengths[i] = length
		}
	case []any:
		if len(v) != args.N {
			return fmt.Errorf("length array must have exactly %d elements, got %d", args.N, len(v))
		}
		totalLength := 0
		for i, l := range v {
			length := int(l.(float64))
			if length < 1 || length > 10000 {
				return fmt.Errorf("parameter 'length[%d]' must be between 1 and 10000, got %d", i, length)
			}
			lengths[i] = length
			totalLength += length
		}
		if totalLength < 1 || totalLength > 10000 {
			return fmt.Errorf("total length of all sequences must be between 1 and 10000, got %d", totalLength)
		}
	default:
		return fmt.Errorf("parameter 'length' must be an integer or array of integers")
	}

	// Normalize mins
	mins := make([]int64, args.N)
	switch v := args.Min.(type) {
	case float64:
		min := int64(v)
		if min < -1e9 || min > 1e9 {
			return fmt.Errorf("parameter 'min' must be between -1e9 and 1e9, got %d", min)
		}
		for i := range mins {
			mins[i] = min
		}
	case []any:
		if len(v) != args.N {
			return fmt.Errorf("min array must have exactly %d elements, got %d", args.N, len(v))
		}
		for i, m := range v {
			min := int64(m.(float64))
			if min < -1e9 || min > 1e9 {
				return fmt.Errorf("parameter 'min[%d]' must be between -1e9 and 1e9, got %d", i, min)
			}
			mins[i] = min
		}
	default:
		return fmt.Errorf("parameter 'min' must be an integer or array of integers")
	}

	// Normalize maxs
	maxs := make([]int64, args.N)
	switch v := args.Max.(type) {
	case float64:
		max := int64(v)
		if max < -1e9 || max > 1e9 {
			return fmt.Errorf("parameter 'max' must be between -1e9 and 1e9, got %d", max)
		}
		for i := range maxs {
			maxs[i] = max
		}
	case []any:
		if len(v) != args.N {
			return fmt.Errorf("max array must have exactly %d elements, got %d", args.N, len(v))
		}
		for i, m := range v {
			max := int64(m.(float64))
			if max < -1e9 || max > 1e9 {
				return fmt.Errorf("parameter 'max[%d]' must be between -1e9 and 1e9, got %d", i, max)
			}
			maxs[i] = max
		}
	default:
		return fmt.Errorf("parameter 'max' must be an integer or array of integers")
	}

	// Normalize replacements (default is true)
	replacements := make([]bool, args.N)

	switch v := args.Replacement.(type) {
	case nil:
		for i := range replacements {
			replacements[i] = true
		}
	case bool:
		for i := range replacements {
			replacements[i] = v
		}
	case []any:
		if len(v) != args.N {
			return fmt.Errorf(
				"replacement array must have exactly %d elements, got %d",
				args.N, len(v),
			)
		}
		for i, x := range v {
			b, ok := x.(bool)
			if !ok {
				return fmt.Errorf(
					"replacement[%d] must be a boolean, got %T",
					i, x,
				)
			}
			replacements[i] = b
		}
	default:
		return fmt.Errorf(
			"parameter 'replacement' must be a boolean or array of booleans, got %T",
			v,
		)
	}

	// Normalize bases
	bases := make([]int, args.N)
	switch v := args.Base.(type) {
	case float64:
		base := int(v)
		if base != 0 && base != 2 && base != 8 && base != 10 && base != 16 {
			return fmt.Errorf("parameter 'base' must be 2, 8, 10, or 16, got %d", base)
		}
		for i := range bases {
			bases[i] = base
		}
	case []any:
		if len(v) != args.N {
			return fmt.Errorf("base array must have exactly %d elements, got %d", args.N, len(v))
		}
		for i, b := range v {
			base := int(b.(float64))
			if base != 0 && base != 2 && base != 8 && base != 10 && base != 16 {
				return fmt.Errorf("parameter 'base[%d]' must be 2, 8, 10, or 16, got %d", i, base)
			}
			bases[i] = base
		}
	case nil:
		for i := range bases {
			bases[i] = 10
		}
	default:
		return fmt.Errorf("parameter 'base' must be an integer or array of integers")
	}

	startTime := time.Now()
	allData := make([][]any, args.N)
	bitsUsed := 0

	for i := 0; i < args.N; i++ {
		length := lengths[i]
		minVal := mins[i]
		maxVal := maxs[i]
		base := bases[i]
		if minVal > maxVal {
			return fmt.Errorf("min[%d] > max[%d] (%d > %d)", i, i, minVal, maxVal)
		}
		domainSize := maxVal - minVal + 1

		var seq []any
		if replacements[i] {
			seq = make([]any, length)
			for j := 0; j < length; j++ {
				num, err := s.generateRandomInt(minVal, maxVal)
				if err != nil {
					log.Printf("Error generating random int: %v", err)
					return fmt.Errorf("failed to generate random number: %w", err)
				}
				if base == 2 {
					seq[j] = fmt.Sprintf("%b", num)
				} else if base == 8 {
					seq[j] = fmt.Sprintf("%o", num)
				} else if base == 16 {
					seq[j] = fmt.Sprintf("%x", num)
				} else {
					seq[j] = num
				}
				bitsUsed += 64
			}
		} else {
			if int64(length) > domainSize {
				return fmt.Errorf("cannot generate %d unique ints in [%d,%d]", length, minVal, maxVal)
			}
			idxs, err := utils.SampleHybrid(int(domainSize), length)
			if err != nil {
				return fmt.Errorf("sampling failed: %w", err)
			}
			seq = make([]any, length)
			for j, idx := range idxs {
				num := minVal + int64(idx)
				if base == 2 {
					seq[j] = fmt.Sprintf("%b", num)
				} else if base == 8 {
					seq[j] = fmt.Sprintf("%o", num)
				} else if base == 16 {
					seq[j] = fmt.Sprintf("%x", num)
				} else {
					seq[j] = num
				}
				bitsUsed += 64
			}
		}
		allData[i] = seq
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           allData,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates true random decimal fractions from a uniform distribution across the [0,1) interval with a user-defined number of decimal places
func (s *RandomService) GenerateDecimalFractions(r *http.Request, args *GenerateDecimalFractionsArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 10000 {
		return fmt.Errorf("parameter 'n' must be between 1 and 10000, got %d", args.N)
	}
	if args.DecimalPlaces < 1 || args.DecimalPlaces > 14 {
		return fmt.Errorf("parameter 'decimalPlaces' must be between 1 and 14, got %d", args.DecimalPlaces)
	}
	if args.Replacement == nil {
		defaultReplacement := true
		args.Replacement = &defaultReplacement
	}

	startTime := time.Now()
	data := make([]float64, 0, args.N)
	bitsUsed := 0

	maxValue := int(sampleInt64(math.Pow10(args.DecimalPlaces)))

	if *args.Replacement {
		for i := 0; i < args.N; i++ {
			num, err := s.generateRandomInt(0, int64(maxValue-1))
			if err != nil {
				log.Printf("Error generating random int: %v", err)
				return fmt.Errorf("failed to generate random number from system source: %w", err)
			}
			fraction := float64(num) / float64(maxValue)
			data = append(data, fraction)
			bitsUsed += 64
		}
	} else {
		if args.N > maxValue {
			return fmt.Errorf("cannot generate %d unique decimal fractions with %d decimal places", args.N, args.DecimalPlaces)
		}

		indices, err := utils.SampleHybrid(maxValue, args.N)
		if err != nil {
			return fmt.Errorf("failed to sample unique indices: %w", err)
		}
		for _, num := range indices {
			fraction := float64(num) / float64(maxValue)
			data = append(data, fraction)
			bitsUsed += 64
		}
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates random numbers from a Gaussian distribution
func (s *RandomService) GenerateGaussians(r *http.Request, args *GenerateGaussiansArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 10000 {
		return fmt.Errorf("parameter 'n' must be between 1 and 10000, got %d", args.N)
	}
	if args.Mean < -1000000 || args.Mean > 1000000 {
		return fmt.Errorf("parameter 'mean' must be between -1000000 and 1000000, got %f", args.Mean)
	}
	if args.StandardDeviation < -1000000 || args.StandardDeviation > 1000000 {
		return fmt.Errorf("parameter 'standardDeviation' must be between -1000000 and 1000000, got %f", args.StandardDeviation)
	}
	if args.SignificantDigits < 2 || args.SignificantDigits > 14 {
		return fmt.Errorf("parameter 'significantDigits' must be between 2 and 14, got %d", args.SignificantDigits)
	}

	startTime := time.Now()
	data := make([]float64, 0, args.N)
	bitsUsed := 0

	multiplier := math.Pow10(args.SignificantDigits)

	for i := 0; i < args.N; i++ {
		u1, err := s.generateRandomInt(1, math.MaxInt64)
		if err != nil {
			log.Printf("Error generating random int: %v", err)
			return fmt.Errorf("failed to generate random number from system source: %w", err)
		}
		u2, err := s.generateRandomInt(1, math.MaxInt64)
		if err != nil {
			log.Printf("Error generating random int: %v", err)
			return fmt.Errorf("failed to generate random number from system source: %w", err)
		}
		bitsUsed += 128

		u1f := float64(u1) / float64(math.MaxInt64)
		u2f := float64(u2) / float64(math.MaxInt64)

		z0 := math.Sqrt(-2.0*math.Log(u1f)) * math.Cos(2.0*math.Pi*u2f)

		gaussian := args.Mean + args.StandardDeviation*z0

		gaussian = math.Round(gaussian*multiplier) / multiplier
		data = append(data, gaussian)
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates random strings
func (s *RandomService) GenerateStrings(r *http.Request, args *GenerateStringsArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 10000 {
		return fmt.Errorf("parameter 'n' must be between 1 and 10000, got %d", args.N)
	}
	if args.Length < 1 || args.Length > 32 {
		return fmt.Errorf("parameter 'length' must be between 1 and 32, got %d", args.Length)
	}
	if len(args.Characters) == 0 || len(args.Characters) > 128 {
		return fmt.Errorf("parameter 'characters' must have length between 1 and 128, got %d", len(args.Characters))
	}
	if args.Replacement == nil {
		defaultReplacement := true
		args.Replacement = &defaultReplacement
	}

	startTime := time.Now()
	data := make([]string, 0, args.N)
	bitsUsed := 0

	bitsPerChar := int(math.Ceil(math.Log2(float64(len(args.Characters)))))

	if *args.Replacement {
		for i := 0; i < args.N; i++ {
			str := make([]byte, args.Length)
			for j := 0; j < args.Length; j++ {
				idx, err := s.generateRandomInt(0, int64(len(args.Characters)-1))
				if err != nil {
					log.Printf("Error generating random int: %v", err)
					return fmt.Errorf("failed to generate random number from system source: %w", err)
				}
				str[j] = args.Characters[idx]
				bitsUsed += bitsPerChar
			}
			data = append(data, string(str))
		}
	} else {
		possible := int64(math.Pow(
			float64(len(args.Characters)),
			float64(args.Length),
		))
		if int64(args.N) > possible {
			return fmt.Errorf(
				"cannot generate %d unique strings of length %d with %d characters",
				args.N, args.Length, len(args.Characters),
			)
		}

		indices, err := utils.SampleHybrid(int(possible), args.N)
		if err != nil {
			return fmt.Errorf("sampling failed: %w", err)
		}
		for _, idx := range indices {
			str := make([]byte, args.Length)
			for j := args.Length - 1; j >= 0; j-- {
				str[j] = args.Characters[idx%len(args.Characters)]
				idx /= len(args.Characters)
			}
			data = append(data, string(str))
			bitsUsed += args.Length * bitsPerChar
		}
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates version 4 random Universally Unique IDentifiers (UUIDs) in accordance with section 4.4 of RFC 4122
func (s *RandomService) GenerateUUIDs(r *http.Request, args *GenerateUUIDsArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 1000 {
		return fmt.Errorf("parameter 'n' must be between 1 and 1000, got %d", args.N)
	}

	startTime := time.Now()
	data := make([]string, 0, args.N)
	bitsUsed := 0

	for i := 0; i < args.N; i++ {
		uuidBytes := make([]byte, 16)
		_, err := rand.Read(uuidBytes)
		if err != nil {
			log.Printf("Error generating random bytes: %v", err)
			return fmt.Errorf("failed to generate random bytes: %w", err)
		}

		uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // Version 4
		uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // Variant RFC4122

		uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
			uuidBytes[0:4],
			uuidBytes[4:6],
			uuidBytes[6:8],
			uuidBytes[8:10],
			uuidBytes[10:16])

		data = append(data, uuid)
		bitsUsed += 128
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}

// generates Binary Large OBjects (BLOBs) containing random data
func (s *RandomService) GenerateBlobs(r *http.Request, args *GenerateBlobsArgs, reply *RandomResult) error {
	if err := s.validateAPIKey(args.APIKey); err != nil {
		return err
	}

	if args.N < 1 || args.N > 100 {
		return fmt.Errorf("parameter 'n' must be between 1 and 100, got %d", args.N)
	}
	if args.Size < 1 || args.Size > 1048576 {
		return fmt.Errorf("parameter 'size' must be between 1 and 1048576, got %d", args.Size)
	}
	if args.Size%8 != 0 {
		return fmt.Errorf("parameter 'size' must be divisible by 8, got %d", args.Size)
	}
	if args.N*args.Size > 1048576 {
		return fmt.Errorf("total size of all blobs must not exceed 1048576 bits, got %d", args.N*args.Size)
	}

	if args.Format != "" && args.Format != "base64" && args.Format != "hex" {
		return fmt.Errorf("parameter 'format' must be either 'base64' or 'hex', got %s", args.Format)
	}

	startTime := time.Now()
	data := make([]string, 0, args.N)
	bitsUsed := 0

	for i := 0; i < args.N; i++ {
		bytesNeeded := args.Size / 8
		blob := make([]byte, bytesNeeded)

		_, err := rand.Read(blob)
		if err != nil {
			log.Printf("Error generating random bytes: %v", err)
			return fmt.Errorf("failed to generate random bytes: %w", err)
		}

		// Format the blob according to the requested format
		var formattedBlob string
		if args.Format == "hex" {
			formattedBlob = fmt.Sprintf("%x", blob)
		} else {
			formattedBlob = base64.StdEncoding.EncodeToString(blob)
		}

		data = append(data, formattedBlob)
		bitsUsed += args.Size
	}

	completionTime := startTime.UTC().Format(time.RFC3339)

	reply.Random = ResultData{
		Data:           data,
		CompletionTime: completionTime,
	}
	reply.BitsUsed = bitsUsed

	return nil
}
