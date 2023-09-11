package clusterautoscaler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"

	"github.com/openshift/rosa/pkg/interactive"
	"github.com/openshift/rosa/pkg/ocm"
)

const (
	balanceSimilarNodeGroupsFlag      = "balance-similar-node-groups"
	skipNodesWithLocalStorageFlag     = "skip-nodes-with-local-storage"
	logVerbosityFlag                  = "log-verbosity"
	maxPodGracePeriodFlag             = "max-pod-grace-period"
	podPriorityThresholdFlag          = "pod-priority-threshold"
	ignoreDaemonsetsUtilizationFlag   = "ignore-daemonsets-utilization"
	maxNodeProvisionTimeFlag          = "max-node-provision-time"
	balancingIgnoredLabelsFlag        = "balancing-ignored-labels"
	maxNodesTotalFlag                 = "max-nodes-total"
	minCoresFlag                      = "min-cores"
	maxCoresFlag                      = "max-cores"
	minMemoryFlag                     = "min-memory"
	maxMemoryFlag                     = "max-memory"
	gpuLimitFlag                      = "gpu-limit"
	scaleDownEnabledFlag              = "scale-down-enabled"
	scaleDownUnneededTimeFlag         = "scale-down-unneeded-time"
	scaleDownUtilizationThresholdFlag = "scale-down-utilization-threshold"
	scaleDownDelayAfterAddFlag        = "scale-down-delay-after-add"
	scaleDownDelayAfterDeleteFlag     = "scale-down-delay-after-delete"
	scaleDownDelayAfterFailureFlag    = "scale-down-delay-after-failure"
)

type AutoscalerArgs struct {
	BalanceSimilarNodeGroups    bool
	SkipNodesWithLocalStorage   bool
	LogVerbosity                int
	MaxPodGracePeriod           int
	PodPriorityThreshold        int
	IgnoreDaemonsetsUtilization bool
	MaxNodeProvisionTime        string
	BalancingIgnoredLabels      []string
	ResourceLimits              ResourceLimits
	ScaleDown                   ScaleDownConfig
}

func IsAutoscalerSetViaCLI(cmd *pflag.FlagSet) bool {
	return cmd.Changed(balanceSimilarNodeGroupsFlag) ||
		cmd.Changed(skipNodesWithLocalStorageFlag) ||
		cmd.Changed(logVerbosityFlag) ||
		cmd.Changed(balancingIgnoredLabelsFlag) ||
		cmd.Changed(ignoreDaemonsetsUtilizationFlag) ||
		cmd.Changed(maxPodGracePeriodFlag) ||
		cmd.Changed(podPriorityThresholdFlag) ||
		cmd.Changed(maxNodeProvisionTimeFlag) ||
		cmd.Changed(maxNodesTotalFlag) ||
		cmd.Changed(minCoresFlag) ||
		cmd.Changed(maxCoresFlag) ||
		cmd.Changed(minMemoryFlag) ||
		cmd.Changed(maxMemoryFlag) ||
		cmd.Changed(gpuLimitFlag) ||
		cmd.Changed(scaleDownEnabledFlag) ||
		cmd.Changed(scaleDownUnneededTimeFlag) ||
		cmd.Changed(scaleDownUtilizationThresholdFlag) ||
		cmd.Changed(scaleDownDelayAfterAddFlag) ||
		cmd.Changed(scaleDownDelayAfterDeleteFlag) ||
		cmd.Changed(scaleDownDelayAfterFailureFlag)
}

type ResourceLimits struct {
	MaxNodesTotal int
	Cores         ResourceRange
	Memory        ResourceRange
	GPULimits     []string
}

type ResourceRange struct {
	Min int
	Max int
}

type ScaleDownConfig struct {
	Enabled              bool
	UnneededTime         string
	UtilizationThreshold float64
	DelayAfterAdd        string
	DelayAfterDelete     string
	DelayAfterFailure    string
}

func AddClusterAutoscalerFlags(cmd *pflag.FlagSet, prefix string) *AutoscalerArgs {
	args := &AutoscalerArgs{}

	cmd.BoolVar(
		&args.BalanceSimilarNodeGroups,
		fmt.Sprintf("%s%s", prefix, balanceSimilarNodeGroupsFlag),
		false,
		"Identify node groups with the same instance type and label set, "+
			"and aim to balance respective sizes of those node groups.",
	)

	cmd.BoolVar(
		&args.SkipNodesWithLocalStorage,
		fmt.Sprintf("%s%s", prefix, skipNodesWithLocalStorageFlag),
		false,
		"If true cluster autoscaler will never delete nodes with pods with local storage, e.g. EmptyDir or HostPath.",
	)

	cmd.IntVar(
		&args.LogVerbosity,
		fmt.Sprintf("%s%s", prefix, logVerbosityFlag),
		1,
		"Autoscaler log level. Default is 1, 4 is a good option when trying to debug the autoscaler.",
	)

	cmd.IntVar(
		&args.MaxPodGracePeriod,
		fmt.Sprintf("%s%s", prefix, maxPodGracePeriodFlag),
		0,
		"Gives pods graceful termination time before scaling down, measured in seconds.",
	)

	cmd.IntVar(
		&args.PodPriorityThreshold,
		fmt.Sprintf("%s%s", prefix, podPriorityThresholdFlag),
		0,
		"The priority that a pod must exceed to cause the cluster autoscaler to deploy additional nodes. "+
			"Expects an integer, can be negative.",
	)

	cmd.BoolVar(
		&args.IgnoreDaemonsetsUtilization,
		fmt.Sprintf("%s%s", prefix, ignoreDaemonsetsUtilizationFlag),
		false,
		"Should cluster-autoscaler ignore DaemonSet pods when calculating resource utilization for scaling down.",
	)

	cmd.StringVar(
		&args.MaxNodeProvisionTime,
		fmt.Sprintf("%s%s", prefix, maxNodeProvisionTimeFlag),
		"",
		"Maximum time cluster-autoscaler waits for node to be provisioned. "+
			"Expects string comprised of an integer and time unit (ns|us|µs|ms|s|m|h), examples: 20m, 1h.",
	)

	cmd.StringSliceVar(
		&args.BalancingIgnoredLabels,
		fmt.Sprintf("%s%s", prefix, balancingIgnoredLabelsFlag),
		nil,
		"A comma-separated list of label keys that cluster autoscaler should ignore when considering node group similarity.",
	)

	// Resource Limits

	cmd.IntVar(
		&args.ResourceLimits.MaxNodesTotal,
		fmt.Sprintf("%s%s", prefix, maxNodesTotalFlag),
		180,
		"Total amount of nodes that can exist in the cluster, including non-scaled nodes.",
	)

	cmd.IntVar(
		&args.ResourceLimits.Cores.Min,
		fmt.Sprintf("%s%s", prefix, minCoresFlag),
		0,
		"Minimum limit for the amount of cores to deploy in the cluster.",
	)

	cmd.IntVar(
		&args.ResourceLimits.Cores.Max,
		fmt.Sprintf("%s%s", prefix, maxCoresFlag),
		100,
		"Maximum limit for the amount of cores to deploy in the cluster.",
	)

	cmd.IntVar(
		&args.ResourceLimits.Memory.Min,
		fmt.Sprintf("%s%s", prefix, minMemoryFlag),
		0,
		"Minimum limit for the amount of memory, in GiB, in the cluster.",
	)

	cmd.IntVar(
		&args.ResourceLimits.Memory.Max,
		fmt.Sprintf("%s%s", prefix, maxMemoryFlag),
		4096,
		"Maximum limit for the amount of memory, in GiB, in the cluster.",
	)

	flag := fmt.Sprintf("%s%s", prefix, gpuLimitFlag)
	cmd.StringArrayVar(
		&args.ResourceLimits.GPULimits,
		flag,
		[]string{},
		fmt.Sprintf(
			"Limit GPUs consumption. It should be comprised of 3 values separated "+
				"with commas: the GPU hardware type, a minimal count for that type "+
				"and a maximal count for that type. This option can be repeated multiple "+
				"times in order to apply multiple restrictions for different GPU types. For example: "+
				"--%[1]s nvidia.com/gpu,0,10 --%[1]s amd.com/gpu,1,5", flag),
	)

	// Scale down Configuration

	cmd.BoolVar(
		&args.ScaleDown.Enabled,
		fmt.Sprintf("%s%s", prefix, scaleDownEnabledFlag),
		false,
		"Should cluster-autoscaler be able to scale down the cluster.",
	)

	cmd.StringVar(
		&args.ScaleDown.UnneededTime,
		fmt.Sprintf("%s%s", prefix, scaleDownUnneededTimeFlag),
		"",
		"Increasing value will make nodes stay up longer, waiting for pods to be scheduled "+
			"while decreasing value will make nodes be deleted sooner.",
	)

	cmd.Float64Var(
		&args.ScaleDown.UtilizationThreshold,
		fmt.Sprintf("%s%s", prefix, scaleDownUtilizationThresholdFlag),
		0.5,
		"Node utilization level, defined as sum of requested resources divided by capacity, "+
			"below which a node can be considered for scale down. Value should be between 0 and 1.",
	)

	cmd.StringVar(
		&args.ScaleDown.DelayAfterAdd,
		fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterAddFlag),
		"",
		"After a scale-up, consider scaling down only after this amount of time.",
	)

	cmd.StringVar(
		&args.ScaleDown.DelayAfterDelete,
		fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterDeleteFlag),
		"",
		"After a scale-down, consider scaling down again only after this amount of time.",
	)

	cmd.StringVar(
		&args.ScaleDown.DelayAfterFailure,
		fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterFailureFlag),
		"",
		"After a failing scale-down, consider scaling down again only after this amount of time.",
	)

	return args
}

func GetAutoscalerOptions(
	cmd *pflag.FlagSet, prefix string, confirmBeforeAllArgs bool, autoscalerArgs *AutoscalerArgs,
) (*AutoscalerArgs, error) {

	var err error
	result := &AutoscalerArgs{}

	result.BalanceSimilarNodeGroups = autoscalerArgs.BalanceSimilarNodeGroups
	result.SkipNodesWithLocalStorage = autoscalerArgs.SkipNodesWithLocalStorage
	result.LogVerbosity = autoscalerArgs.LogVerbosity
	result.MaxPodGracePeriod = autoscalerArgs.MaxPodGracePeriod
	result.PodPriorityThreshold = autoscalerArgs.PodPriorityThreshold
	result.IgnoreDaemonsetsUtilization = autoscalerArgs.IgnoreDaemonsetsUtilization
	result.MaxNodeProvisionTime = autoscalerArgs.MaxNodeProvisionTime
	result.BalancingIgnoredLabels = autoscalerArgs.BalancingIgnoredLabels
	result.ResourceLimits.MaxNodesTotal = autoscalerArgs.ResourceLimits.MaxNodesTotal
	result.ResourceLimits.Cores.Min = autoscalerArgs.ResourceLimits.Cores.Min
	result.ResourceLimits.Cores.Max = autoscalerArgs.ResourceLimits.Cores.Max
	result.ResourceLimits.Memory.Min = autoscalerArgs.ResourceLimits.Memory.Min
	result.ResourceLimits.Memory.Max = autoscalerArgs.ResourceLimits.Memory.Max
	result.ResourceLimits.GPULimits = append(result.ResourceLimits.GPULimits, autoscalerArgs.ResourceLimits.GPULimits...)
	result.ScaleDown.Enabled = autoscalerArgs.ScaleDown.Enabled
	result.ScaleDown.UnneededTime = autoscalerArgs.ScaleDown.UnneededTime
	result.ScaleDown.UtilizationThreshold = autoscalerArgs.ScaleDown.UtilizationThreshold
	result.ScaleDown.DelayAfterAdd = autoscalerArgs.ScaleDown.DelayAfterAdd
	result.ScaleDown.DelayAfterDelete = autoscalerArgs.ScaleDown.DelayAfterDelete
	result.ScaleDown.DelayAfterFailure = autoscalerArgs.ScaleDown.DelayAfterFailure

	isBalanceSimilarNodeGroupsSet := cmd.Changed(balanceSimilarNodeGroupsFlag)
	isAutoscalerSkipNodesWithLocalStorageSet := cmd.Changed(skipNodesWithLocalStorageFlag)
	isAutoscalerLogVerbositySet := cmd.Changed(logVerbosityFlag)
	isAutoscalerBalancingIgnoredLabelsSet := cmd.Changed(balancingIgnoredLabelsFlag)
	isAutoscalerIgnoreDaemonsetsUtilizationSet := cmd.Changed(ignoreDaemonsetsUtilizationFlag)
	isAutoscalerMaxPodGracePeriodSet := cmd.Changed(maxPodGracePeriodFlag)
	isAutoscalerPodPriorityThresholdSet := cmd.Changed(podPriorityThresholdFlag)
	isAutoscalerMaxNodeProvisionTimeSet := cmd.Changed(maxNodeProvisionTimeFlag)
	isMaxNodesTotalSet := cmd.Changed(maxNodesTotalFlag)
	isMinCoresSet := cmd.Changed(minCoresFlag)
	isMaxCoresSet := cmd.Changed(maxCoresFlag)
	isMinMemorySet := cmd.Changed(minMemoryFlag)
	isMaxMemorySet := cmd.Changed(maxMemoryFlag)
	isGPULimitsSet := cmd.Changed(gpuLimitFlag)
	isScaleDownEnabledSet := cmd.Changed(scaleDownEnabledFlag)
	isScaleDownUnneededTimeSet := cmd.Changed(scaleDownUnneededTimeFlag)
	isScaleDownUtilizationThresholdSet := cmd.Changed(scaleDownUtilizationThresholdFlag)
	isScaleDownDelayAfterAddSet := cmd.Changed(scaleDownDelayAfterAddFlag)
	isScaleDownDelayAfterDeleteSet := cmd.Changed(scaleDownDelayAfterDeleteFlag)
	isScaleDownDelayAfterFailureSet := cmd.Changed(scaleDownDelayAfterFailureFlag)

	isClusterAutoscalerSet := IsAutoscalerSetViaCLI(cmd)

	if confirmBeforeAllArgs && !isClusterAutoscalerSet && interactive.Enabled() {
		isClusterAutoscalerSet, err = interactive.GetBool(interactive.Input{
			Question: "Configure cluster-autoscaler",
			Help:     "Set cluster-wide autoscaling configurations",
			Default:  false,
			Required: false,
		})
		if err != nil {
			return nil, err
		}

		if !isClusterAutoscalerSet {
			// User is not interested providing any cluster-autoscaler parameter
			return nil, nil
		}
	}

	if interactive.Enabled() && !isBalanceSimilarNodeGroupsSet {
		result.BalanceSimilarNodeGroups, err = interactive.GetBool(interactive.Input{
			Question: "Balance similar node groups",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, balanceSimilarNodeGroupsFlag)).Usage,
			Default:  result.BalanceSimilarNodeGroups,
			Required: false,
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isAutoscalerSkipNodesWithLocalStorageSet {
		result.SkipNodesWithLocalStorage, err = interactive.GetBool(interactive.Input{
			Question: "Skip nodes with local storage",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, skipNodesWithLocalStorageFlag)).Usage,
			Default:  result.SkipNodesWithLocalStorage,
			Required: false,
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isAutoscalerLogVerbositySet {
		result.LogVerbosity, err = interactive.GetInt(interactive.Input{
			Question: "Log verbosity",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, logVerbosityFlag)).Usage,
			Default:  result.LogVerbosity,
			Required: false,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.IntValidator(result.LogVerbosity); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isAutoscalerBalancingIgnoredLabelsSet {
		balancingIgnoredLabels, err := interactive.GetString(interactive.Input{
			Question: "Labels that cluster autoscaler should ignore when considering node group similarity",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, balancingIgnoredLabelsFlag)).Usage,
			Default:  strings.Join(result.BalancingIgnoredLabels, ","),
			Required: false,
			Validators: []interactive.Validator{
				ocm.ValidateBalancingIgnoredLabels,
			},
		})
		if err != nil {
			return nil, err
		}

		if balancingIgnoredLabels != "" {
			result.BalancingIgnoredLabels = strings.Split(balancingIgnoredLabels, ",")
		}
	}

	if interactive.Enabled() && !isAutoscalerIgnoreDaemonsetsUtilizationSet {
		result.IgnoreDaemonsetsUtilization, err = interactive.GetBool(interactive.Input{
			Question: "Ignore daemonsets utilization",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, ignoreDaemonsetsUtilizationFlag)).Usage,
			Default:  result.IgnoreDaemonsetsUtilization,
			Required: false,
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isAutoscalerMaxNodeProvisionTimeSet {
		result.MaxNodeProvisionTime, err = interactive.GetString(interactive.Input{
			Question: "Maximum node provision time",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, maxNodeProvisionTimeFlag)).Usage,
			Required: false,
			Default:  result.MaxNodeProvisionTime,
			Validators: []interactive.Validator{
				ocm.PositiveDurationStringValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.PositiveDurationStringValidator(result.MaxNodeProvisionTime); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isAutoscalerMaxPodGracePeriodSet {
		result.MaxPodGracePeriod, err = interactive.GetInt(interactive.Input{
			Question: "Maximum pod grace period",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, maxPodGracePeriodFlag)).Usage,
			Required: false,
			Default:  result.MaxPodGracePeriod,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isAutoscalerPodPriorityThresholdSet {
		result.PodPriorityThreshold, err = interactive.GetInt(interactive.Input{
			Question: "Pod priority threshold",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, podPriorityThresholdFlag)).Usage,
			Required: false,
			Default:  result.PodPriorityThreshold,
			Validators: []interactive.Validator{
				ocm.IntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isMaxNodesTotalSet {
		result.ResourceLimits.MaxNodesTotal, err = interactive.GetInt(interactive.Input{
			Question: "Maximum amount of nodes in the cluster",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, maxNodesTotalFlag)).Usage,
			Required: false,
			Default:  result.ResourceLimits.MaxNodesTotal,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.NonNegativeIntValidator(result.ResourceLimits.MaxNodesTotal); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isMinCoresSet {
		result.ResourceLimits.Cores.Min, err = interactive.GetInt(interactive.Input{
			Question: "Minimum number of cores to deploy in cluster",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, minCoresFlag)).Usage,
			Required: false,
			Default:  result.ResourceLimits.Cores.Min,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err = ocm.NonNegativeIntValidator(result.ResourceLimits.Cores.Min); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isMaxCoresSet {
		result.ResourceLimits.Cores.Max, err = interactive.GetInt(interactive.Input{
			Question: "Maximum number of cores to deploy in cluster",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, maxCoresFlag)).Usage,
			Required: false,
			Default:  result.ResourceLimits.Cores.Max,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
				getValidMaxRangeValidator(result.ResourceLimits.Cores.Min),
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.NonNegativeIntValidator(result.ResourceLimits.Cores.Max); err != nil {
		return nil, err
	}
	if err := getValidMaxRangeValidator(result.ResourceLimits.Cores.Min)(result.ResourceLimits.Cores.Max); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isMinMemorySet {
		result.ResourceLimits.Memory.Min, err = interactive.GetInt(interactive.Input{
			Question: "Minimum amount of memory, in GiB, in the cluster",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, minMemoryFlag)).Usage,
			Required: false,
			Default:  result.ResourceLimits.Memory.Min,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.NonNegativeIntValidator(result.ResourceLimits.Memory.Min); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isMaxMemorySet {
		result.ResourceLimits.Memory.Max, err = interactive.GetInt(interactive.Input{
			Question: "Maximum amount of memory, in GiB, in the cluster",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, maxMemoryFlag)).Usage,
			Required: false,
			Default:  result.ResourceLimits.Memory.Max,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
				getValidMaxRangeValidator(result.ResourceLimits.Memory.Min),
			},
		})
		if err != nil {
			return nil, err
		}
	}
	if err := ocm.NonNegativeIntValidator(result.ResourceLimits.Memory.Max); err != nil {
		return nil, err
	}
	if err := getValidMaxRangeValidator(result.ResourceLimits.Memory.Min)(result.ResourceLimits.Memory.Max); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isGPULimitsSet {
		gpuLimitsCount, err := interactive.GetInt(interactive.Input{
			Question: "Enter the number of GPU limitations you wish to set",
			Help: "This allows setting a limiting range of the count of GPU resources " +
				"that will be used. Each limitation is per hardware type",
			Default:  0,
			Required: false,
			Validators: []interactive.Validator{
				ocm.NonNegativeIntValidator,
			},
		})

		if err != nil {
			return nil, err
		}

		for i := 1; i <= gpuLimitsCount; i++ {
			gpuLimitType, err := interactive.GetString(interactive.Input{
				Question: fmt.Sprintf("%d. Enter the type of desired GPU limitation", i),
				Help:     "E.g.: nvidia.com/gpu, amd.com/gpu",
				Required: true,
			})

			if err != nil {
				return nil, err
			}

			gpuLimitMin, err := interactive.GetInt(interactive.Input{
				Question: fmt.Sprintf("%d. Enter minimum number of GPUS of type '%s' to deploy in the cluster.", i, gpuLimitType),
				Help: "An integer stating the minimum number of GPUs of the given type to deploy in the cluster. " +
					"Must always be smaller than or equal to the maximal value.",
				Validators: []interactive.Validator{
					ocm.NonNegativeIntValidator,
				},
			})

			if err != nil {
				return nil, err
			}

			gpuLimitMax, err := interactive.GetInt(interactive.Input{
				Question: fmt.Sprintf("%d. Enter maximum number of GPUS of type '%s' to deploy in the cluster.", i, gpuLimitType),
				Help: "An integer stating the maximum number of GPUs of the given type to deploy in the cluster. " +
					"Must always be smaller than or equal to the maximal value.",
				Validators: []interactive.Validator{
					ocm.NonNegativeIntValidator,
					getValidMaxRangeValidator(gpuLimitMin),
				},
			})

			if err != nil {
				return nil, err
			}

			result.ResourceLimits.GPULimits = append(result.ResourceLimits.GPULimits,
				fmt.Sprintf("%s,%d,%d", gpuLimitType, gpuLimitMin, gpuLimitMax))
		}
	}

	// scale-down configs

	if interactive.Enabled() && !isScaleDownEnabledSet {
		result.ScaleDown.Enabled, err = interactive.GetBool(interactive.Input{
			Question: "Should scale-down be enabled",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownEnabledFlag)).Usage,
			Default:  result.ScaleDown.Enabled,
			Required: false,
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isScaleDownUnneededTimeSet {
		result.ScaleDown.UnneededTime, err = interactive.GetString(interactive.Input{
			Question: "How long a node should be unneeded before it is eligible for scale down",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownUnneededTimeFlag)).Usage,
			Default:  result.ScaleDown.UnneededTime,
			Required: false,
			Validators: []interactive.Validator{
				ocm.PositiveDurationStringValidator,
			},
		})
		if err != nil {
			return nil, err
		}

		if err := ocm.PositiveDurationStringValidator(result.ScaleDown.UnneededTime); err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isScaleDownUtilizationThresholdSet {
		result.ScaleDown.UtilizationThreshold, err = interactive.GetFloat(interactive.Input{
			Question: "Node utilization threshold",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownUtilizationThresholdFlag)).Usage,
			Default:  result.ScaleDown.UtilizationThreshold,
			Required: false,
			Validators: []interactive.Validator{
				ocm.PercentageValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isScaleDownDelayAfterAddSet {
		result.ScaleDown.DelayAfterAdd, err = interactive.GetString(interactive.Input{
			Question: "How long after scale up should scale down evaluation resume",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterAddFlag)).Usage,
			Default:  result.ScaleDown.DelayAfterAdd,
			Required: false,
			Validators: []interactive.Validator{
				ocm.PositiveDurationStringValidator,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	if err := ocm.PositiveDurationStringValidator(result.ScaleDown.DelayAfterAdd); err != nil {
		return nil, err
	}

	if interactive.Enabled() && !isScaleDownDelayAfterDeleteSet {
		result.ScaleDown.DelayAfterDelete, err = interactive.GetString(interactive.Input{
			Question: "How long after node deletion should scale down evaluation resume",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterDeleteFlag)).Usage,
			Default:  result.ScaleDown.DelayAfterDelete,
			Required: false,
			Validators: []interactive.Validator{
				ocm.PositiveDurationStringValidator,
			},
		})
		if err != nil {
			return nil, err
		}

		if err := ocm.PositiveDurationStringValidator(result.ScaleDown.DelayAfterDelete); err != nil {
			return nil, err
		}
	}

	if interactive.Enabled() && !isScaleDownDelayAfterFailureSet {
		result.ScaleDown.DelayAfterFailure, err = interactive.GetString(interactive.Input{
			Question: "How long after node deletion failure should scale down evaluation resume.",
			Help:     cmd.Lookup(fmt.Sprintf("%s%s", prefix, scaleDownDelayAfterFailureFlag)).Usage,
			Default:  result.ScaleDown.DelayAfterFailure,
			Required: false,
			Validators: []interactive.Validator{
				ocm.PositiveDurationStringValidator,
			},
		})
		if err != nil {
			return nil, err
		}

		if err := ocm.PositiveDurationStringValidator(result.ScaleDown.DelayAfterFailure); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func BuildAutoscalerOptions(spec *ocm.AutoscalerConfig, prefix string) string {
	if spec == nil {
		return ""
	}

	command := ""

	if spec.BalanceSimilarNodeGroups {
		command += fmt.Sprintf(" --%s%s", prefix, balanceSimilarNodeGroupsFlag)
	}

	if spec.SkipNodesWithLocalStorage {
		command += fmt.Sprintf(" --%s%s", prefix, skipNodesWithLocalStorageFlag)
	}

	command += fmt.Sprintf(" --%s%s %d", prefix, logVerbosityFlag, spec.LogVerbosity)

	command += fmt.Sprintf(" --%s%s %d", prefix, maxPodGracePeriodFlag, spec.MaxPodGracePeriod)

	command += fmt.Sprintf(" --%s%s %d",
		prefix, podPriorityThresholdFlag, spec.PodPriorityThreshold)

	if spec.BalanceSimilarNodeGroups {
		command += fmt.Sprintf(" --%s%s", prefix, ignoreDaemonsetsUtilizationFlag)
	}

	if spec.MaxNodeProvisionTime != "" {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, maxNodeProvisionTimeFlag, spec.MaxNodeProvisionTime)
	}

	if len(spec.BalancingIgnoredLabels) > 0 {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, balancingIgnoredLabelsFlag,
			strings.Join(spec.BalancingIgnoredLabels, ","))
	}

	command += fmt.Sprintf(" --%s%s %d", prefix, maxNodesTotalFlag, spec.ResourceLimits.MaxNodesTotal)

	command += fmt.Sprintf(" --%s%s %d", prefix, minCoresFlag, spec.ResourceLimits.Cores.Min)
	command += fmt.Sprintf(" --%s%s %d", prefix, maxCoresFlag, spec.ResourceLimits.Cores.Max)

	command += fmt.Sprintf(" --%s%s %d", prefix, minMemoryFlag, spec.ResourceLimits.Memory.Min)
	command += fmt.Sprintf(" --%s%s %d", prefix, maxMemoryFlag, spec.ResourceLimits.Memory.Max)

	for _, gpuLimit := range spec.ResourceLimits.GPULimits {
		command += fmt.Sprintf(" --%s%s %s,%d,%d", prefix, gpuLimitFlag,
			gpuLimit.Type, gpuLimit.Range.Min, gpuLimit.Range.Max)
	}

	if spec.ScaleDown.Enabled {
		command += fmt.Sprintf(" --%s%s", prefix, scaleDownEnabledFlag)
	}

	if spec.ScaleDown.UnneededTime != "" {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, scaleDownUnneededTimeFlag, spec.ScaleDown.UnneededTime)
	}

	command += fmt.Sprintf(" --%s%s %f",
		prefix, scaleDownUtilizationThresholdFlag, spec.ScaleDown.UtilizationThreshold)

	if spec.ScaleDown.DelayAfterAdd != "" {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, scaleDownDelayAfterAddFlag, spec.ScaleDown.DelayAfterAdd)
	}

	if spec.ScaleDown.DelayAfterDelete != "" {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, scaleDownDelayAfterDeleteFlag, spec.ScaleDown.DelayAfterDelete)
	}

	if spec.ScaleDown.DelayAfterFailure != "" {
		command += fmt.Sprintf(" --%s%s %s",
			prefix, scaleDownDelayAfterFailureFlag, spec.ScaleDown.DelayAfterFailure)
	}

	return command
}

func CreateAutoscalerConfig(args *AutoscalerArgs) (*ocm.AutoscalerConfig, error) {
	gpuLimits := []ocm.GPULimit{}
	for _, gpuLimit := range args.ResourceLimits.GPULimits {
		parameters := strings.Split(gpuLimit, ",")
		if len(parameters) != 3 {
			return nil, fmt.Errorf("GPU limitation '%s' does not have 3 entries split by a comma", gpuLimit)
		}
		gpuLimitMin, err := strconv.Atoi(parameters[1])
		if err != nil {
			return nil, fmt.Errorf("Failed parsing '%s' into an integer: %s", parameters[1], err)
		}
		gpuLimitMax, err := strconv.Atoi(parameters[2])
		if err != nil {
			return nil, fmt.Errorf("Failed parsing '%s' into an integer: %s", parameters[2], err)
		}

		gpuLimits = append(gpuLimits,
			ocm.GPULimit{
				Type: parameters[0],
				Range: ocm.ResourceRange{
					Min: gpuLimitMin,
					Max: gpuLimitMax,
				},
			})
	}

	return &ocm.AutoscalerConfig{
		BalanceSimilarNodeGroups:    args.BalanceSimilarNodeGroups,
		SkipNodesWithLocalStorage:   args.SkipNodesWithLocalStorage,
		LogVerbosity:                args.LogVerbosity,
		MaxPodGracePeriod:           args.MaxPodGracePeriod,
		BalancingIgnoredLabels:      args.BalancingIgnoredLabels,
		IgnoreDaemonsetsUtilization: args.IgnoreDaemonsetsUtilization,
		MaxNodeProvisionTime:        args.MaxNodeProvisionTime,
		PodPriorityThreshold:        args.PodPriorityThreshold,
		ResourceLimits: ocm.ResourceLimits{
			MaxNodesTotal: args.ResourceLimits.MaxNodesTotal,
			Cores: ocm.ResourceRange{
				Min: args.ResourceLimits.Cores.Min,
				Max: args.ResourceLimits.Cores.Max,
			},
			Memory: ocm.ResourceRange{
				Min: args.ResourceLimits.Memory.Min,
				Max: args.ResourceLimits.Memory.Max,
			},
			GPULimits: gpuLimits,
		},
		ScaleDown: ocm.ScaleDownConfig{
			Enabled:              args.ScaleDown.Enabled,
			UnneededTime:         args.ScaleDown.UnneededTime,
			UtilizationThreshold: args.ScaleDown.UtilizationThreshold,
			DelayAfterAdd:        args.ScaleDown.DelayAfterAdd,
			DelayAfterDelete:     args.ScaleDown.DelayAfterDelete,
			DelayAfterFailure:    args.ScaleDown.DelayAfterFailure,
		},
	}, nil
}

// getValidMaxRangeValidator returns a validator function that asserts a given
// number is greater than or equal to a fixed minimal number.
func getValidMaxRangeValidator(min int) func(interface{}) error {
	return func(val interface{}) error {
		if val == "" { // Allowing optional inputs
			return nil
		}

		max, err := strconv.Atoi(fmt.Sprintf("%v", val))
		if err != nil {
			return fmt.Errorf("Failed parsing '%v' to an integer number.", val)
		}

		if max < min {
			return fmt.Errorf("max value must be greater or equal than min value %d.", min)
		}

		return nil
	}
}
