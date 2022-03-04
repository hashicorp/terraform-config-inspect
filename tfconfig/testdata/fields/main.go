package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	validations := map[string]string{}
	var levelUp []string
	str := "^({((?i)\\\"private\\\":(\\{(\\\"ami_release_version\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"ami_type\\\":\\\"(AL2_x86_64|AL2_x86_64_GPU|AL2_ARM_64|CUSTOM|BOTTLEROCKET_ARM_64|BOTTLEROCKET_x86_64)\\\"(,)?)?(\\\"capacity_type\\\":\\\"(SPOT|ON_DEMAND)\\\"(,)?)?(\\\"disk_size\\\":([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])(,)?)?(\\\"ec2_ssh_key\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"force_update_version\\\":((?i)(true|false))(,)?)?(\\\"iam_role_arn\\\":\\\"((arn:aws:iam::[0-9]{1,}:role\\/)(.+))?\\\"(,)?)?(\\\"instance_types\\\":\\[(\\\"(t2|m3|m4|m5|c4|c5|i3|r3|r4|x1|p2|p3|r5).(([2-9]|[1-2][0-9])?((x)?large|medium))\\\"(,)?)*\\](,)?)?(\\\"max_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"max_unavailable_percentage\\\":([0-9]|[1-9][0-9]|100)(,)?)?(\\\"min_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"subnets\\\":\\[(\\\"subnet-[a-zA-Z0-9]{1,}\\\"(,)?){1,}\\]((,)?)?)?(\\\"timeouts\\\":\\{(\\\"create\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"delete\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"update\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(,)?\\}(,)?)?(\\\"version\\\":(\\\"(1.20|1.21)\\\"|null))?(,)?\\})\\\"private-edge\\\":(\\{(\\\"ami_release_version\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"ami_type\\\":\\\"(AL2_x86_64|AL2_x86_64_GPU|AL2_ARM_64|CUSTOM|BOTTLEROCKET_ARM_64|BOTTLEROCKET_x86_64)\\\"(,)?)?(\\\"capacity_type\\\":\\\"(SPOT|ON_DEMAND)\\\"(,)?)?(\\\"disk_size\\\":([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])(,)?)?(\\\"ec2_ssh_key\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"force_update_version\\\":((?i)(true|false))(,)?)?(\\\"iam_role_arn\\\":\\\"((arn:aws:iam::[0-9]{1,}:role\\/)(.+))?\\\"(,)?)?(\\\"instance_types\\\":\\[(\\\"(t2|m3|m4|m5|c4|c5|i3|r3|r4|x1|p2|p3|r5).(([2-9]|[1-2][0-9])?((x)?large|medium))\\\"(,)?)*\\](,)?)?(\\\"max_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"max_unavailable_percentage\\\":([0-9]|[1-9][0-9]|100)(,)?)?(\\\"min_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"subnets\\\":\\[(\\\"subnet-[a-zA-Z0-9]{1,}\\\"(,)?){1,}\\](,)?)?(\\\"timeouts\\\":\\{(\\\"create\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"delete\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"update\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(,)?\\}(,)?)?(\\\"version\\\":(\\\"(1.20|1.21)\\\"|null))?(,)?\\})\\\"tools\\\":(\\{(\\\"ami_release_version\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"ami_type\\\":\\\"(AL2_x86_64|AL2_x86_64_GPU|AL2_ARM_64|CUSTOM|BOTTLEROCKET_ARM_64|BOTTLEROCKET_x86_64)\\\"(,)?)?(\\\"capacity_type\\\":\\\"(SPOT|ON_DEMAND)\\\"(,)?)?(\\\"disk_size\\\":([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])(,)?)?(\\\"ec2_ssh_key\\\":(\\\"(([a-zA-Z0-9_.-]){1,64})\\\"|null)(,)?)?(\\\"force_update_version\\\":((?i)(true|false))(,)?)?(\\\"iam_role_arn\\\":\\\"((arn:aws:iam::[0-9]{1,}:role\\/)(.+))?\\\"(,)?)?(\\\"instance_types\\\":\\[(\\\"(t2|m3|m4|m5|c4|c5|i3|r3|r4|x1|p2|p3|r5).(([2-9]|[1-2][0-9])?((x)?large|medium))\\\"(,)?)*\\](,)?)?(\\\"max_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"max_unavailable_percentage\\\":([0-9]|[1-9][0-9]|100)(,)?)?(\\\"min_capacity\\\":([1-9]|[1-9][0-9])(,)?)?(\\\"subnets\\\":\\[(\\\"subnet-[a-zA-Z0-9]{1,}\\\"(,)?){1,}\\](,)?)?(\\\"timeouts\\\":\\{(\\\"create\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"delete\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(\\\"update\\\":\\\"(([0-9]){1,3}(s|m))\\\"(,)?)?(,)?\\}(,)?)?(\\\"version\\\":(\\\"(1.20|1.21)\\\"|null))?(,)?\\})})?)$"
	//str := "^\\[(\\{\\\"key_arn\\\":\\\"arn:aws:kms:(.+):key\\/([0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12})\\\"(,)?\\\"resources\\\":\\[\\\"[a-zA-Z0-9_.-]{1,}\\\"\\]\\})?\\]$"
	//str := "^(\\{((\\\"port\\\":(([0-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-5][0-9][0-9][0-9][0-9]|6[0-4][0-9][0-9][0-9]|65[0-4][0-9][0-9]|655[0-2][0-9]|6553[0-4]))(,)?)\\\"type\\\":\\\"(LoadBalancer|ClusterIP|NodePort)\\\")?\\})?$"
	//str := "^\\[({((?i)\\\"groups\\\":\\[(.+)\\](,)?\\\"rolearn\\\":\\\"(arn:aws:iam::[0-9]{1,}:role\\/)(.+)?\\\"(,)?\\\"username\\\":\\\"(.+)\\\"})?\\])$"
	//str := "^\\[({((?i)\\\"groups\\\":\\[(.+)\\](,)?\\\"userarn\\\":\\\"(arn:aws:iam::[0-9]{1,}:user\\/)(.+)?\\\"(,)?\\\"username\\\":\\\"(.+)\\\"})?\\])$"
	//str := "^(\\{(\\\"failureThreshold\\\":([1-9]|[1-9][0-9])(,)?\\\"initialDelaySeconds\\\":([1-9]|[1-9][0-9]|[1-9][0-9][0-9])(,)?\\\"periodSeconds\\\":([1-9]|[1-9][0-9]))\\})?$"
	strSplit := strings.Split(str, "(,)?")
	for y := 0; y < len(strSplit); y++ {
		fName := ""
		for {
			if -1 != strings.Index(strSplit[y], "\\{") {
				strSplitint := strings.Split(strSplit[y], "\\{")
				levels := strings.SplitAfter(strSplitint[0], "\"")
				if len(levels) > 1 {
					levelUp = append(levelUp, strings.Replace(levels[1], "\\\"", "", -1))
				}
				strSplit[y] = strSplitint[1]
				continue
			}
			break
		}
		for {
			if -1 != strings.Index(strSplit[y], "\\}") {
				strSplitint := strings.Split(strSplit[y], "\\}")
				strSplit[y] = strSplitint[0]
				if len(levelUp) > 0 {
					levelUp = levelUp[:len(levelUp)-1]
				}
				continue
			}
			break
		}

		//string
		lastSplit := strings.SplitAfter(strSplit[y], "\"")
		for u := 0; u < len(lastSplit); u++ {
			if len(lastSplit) < 4 {
				break
			}
			_, err := regexp.Compile("^" + strings.Replace(lastSplit[3], "\"", "", -1) + "$")
			if nil == err {
				if len(levelUp) > 0 {
					fName = strings.Join(levelUp, ":") + ":" + strings.Replace(lastSplit[1], "\\\"", "", -1)
				} else {
					fName = strings.Replace(lastSplit[1], "\\\"", "", -1)
				}
				validations[fName] = "^" + strings.Replace(lastSplit[3], "\\\"", "", -1) + "$"
				break
			}
		}
		if len(lastSplit) > 1 && validations[strings.Replace(lastSplit[1], "\\\"", "", -1)] != "" {
			continue
		}
		//fmt.Println(y)
		//number
		numbSplit := strings.SplitAfter(strSplit[y], ":")
		strSplit := strings.SplitAfter(numbSplit[0], "\"")
		if len(numbSplit) < 2 || len(strSplit) < 2 {
			continue
		}
		_, err := regexp.Compile("^" + strings.Replace(numbSplit[1], "\"", "", -1) + "$")
		if nil == err {
			if len(levelUp) > 0 {
				fName = strings.Join(levelUp, ":") + ":" + strings.Replace(lastSplit[1], "\\\"", "", -1)
			} else {
				fName = strings.Replace(strSplit[1], "\\\"", "", -1)
			}
			validations[fName] = "^" + strings.Replace(numbSplit[1], "\\\"", "", -1) + "$"
			continue
		}
	}
	fmt.Printf("%v", validations)
	/*for {
		str = tfconfig.Between(str, "((", "))")

		_, err := regexp.Compile("^" + str + "$")
		if nil == err {
			fmt.Printf("%s", str)
		}

		if str == "" {
			break
		}
	}*/
}
