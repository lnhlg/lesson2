package lesson2

import (
	"fmt"
	"strconv"
	"strings"
)

//SubdomainVisits: 子域名访问计数
/*parameter
cpdomains: 形如["9001 discuss.leetcode.com"]的字符串数组
return: 形如["9001 leetcode.com","9001 discuss.leetcode.com","9001 com"]的字符串数组
*/
func SubdomainVisits(cpdomains []string) []string {

	domainCntMap := make(map[string]int)

	for i := range cpdomains {

		domainCnts := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(domainCnts[0])
		index := 0

		for index >= 0 {

			domainCntMap[domainCnts[1]] += count
			index = strings.Index(domainCnts[1], ".")
			domainCnts[1] = domainCnts[1][index+1:]
		}
	}

	result := make([]string, len(domainCntMap))
	index := 0
	for k, v := range domainCntMap {
		result[index] = fmt.Sprintf("%d %s", v, k)
		index++
	}

	return result
}
