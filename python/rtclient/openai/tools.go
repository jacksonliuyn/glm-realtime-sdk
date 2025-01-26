package oai

import "gitlab.glm.ai/ai-search/ai-e2e-chat/common/constant"

type Tool struct {
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  ToolParameters `json:"parameters"`
}

type ToolParameters struct {
	Type       string                  `json:"type"`
	Properties map[string]ToolProperty `json:"properties"`
	Required   []string                `json:"required"`
}

type ToolProperty struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func (t *Tool) ToEngineFormat() map[string]any {
	r := map[string]any{
		"type": t.Type,
		"function": map[string]any{
			"name":        t.Name,
			"description": t.Description,
			"parameters":  t.Parameters,
		},
	}

	return r
}

func NewAutoSearchTool() Tool {
	return Tool{
		Type: "function",
		Name: constant.ToolCallNameSearchEngine,
		Description: `多功能网络搜索工具，旨在检索互联网上的实时、准确和全面的信息。请在以下场景中策略性地使用此工具：
		1. 信息收集
		- 获取当前事件和最新新闻
		- 检索有关人员、组织和技术的最新事实
		- 收集复杂主题的背景信息
		2. 研究支持
		- 查找专家意见和最新研究
		- 验证声明和交叉引用信息
		- 探索某个主题的多种观点
		3. 上下文查询
		- 解决模棱两可或时间敏感的问题
		- 获得精确的定义和解释
		- 发现特定领域的最新发展
		关键使用指南：
		- 制定精确、有针对性的搜索查询
		- 使用特定关键字来提高结果相关性`,
		Parameters: ToolParameters{
			Type: "object",
			Properties: map[string]ToolProperty{
				"q": {
					Type:        "string",
					Description: "搜索查询",
				},
			},
		},
	}
}
