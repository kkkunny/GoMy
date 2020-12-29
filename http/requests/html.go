package requests

import (
	"golang.org/x/net/html"
	"regexp"
)

// html节点
type HtmlNode struct {
	Node *html.Node
}
// 分割xpath语法
func (n HtmlNode) splitXpathString(xpath string)(string, string){
	ei := 2
	for _, char := range xpath[2:]{
		if char != '/'{
			ei++
		}else{
			break
		}
	}
	this_xpath := xpath[:ei]
	childs_xpath := xpath[ei:]
	return this_xpath, childs_xpath
}
// 处理xpath语法(单)
func (n HtmlNode) procXpathString(xpath string, root bool)*HtmlNode {
	if xpath != ""{
		re := regexp.MustCompile(`(.+)\[@(.+)="(.+)"`)
		xpath_results := re.FindStringSubmatch(xpath)
		var node *HtmlNode
		if len(xpath_results) > 0{
			if root{
				node = n.getNodeByAttrFromChilds(xpath_results[1], xpath_results[2], xpath_results[3])
			}else{
				node = n.GetNodeByAttr(xpath_results[1], xpath_results[2], xpath_results[3])
			}
			return node
		}
		if root{
			node = n.getNodeFromChilds(xpath)
		}else{
			node = n.GetNode(xpath)
		}
		return node
	}
	return nil
}
// 处理xpath语法(多)
func (n HtmlNode) procMultiXpathString(xpath string, root bool)[]*HtmlNode {
	if xpath != ""{
		re := regexp.MustCompile(`(.+)\[@(.+)="(.+)"`)
		xpath_results := re.FindStringSubmatch(xpath)
		var nodes []*HtmlNode
		if len(xpath_results) > 0{
			if root{
				nodes = n.getNodesByAttrFromChilds(xpath_results[1], xpath_results[2], xpath_results[3])
			}else{
				nodes = n.GetNodesByAttr(xpath_results[1], xpath_results[2], xpath_results[3])
			}
			return nodes
		}
		if root{
			nodes = n.getNodesFromChilds(xpath)
		}else{
			nodes = n.GetNodes(xpath)
		}
		return nodes
	}
	return []*HtmlNode{}
}
// 在子节点中获取首个节点
func (n HtmlNode) getNodeFromChilds(label string)*HtmlNode {
	// 检索自身子节点
	childs := n.GetChilds()
	for _, child := range childs{
		if child.Node.Data == label{
			return child
		}
	}
	return nil
}
// 在子节点中获取多个节点
func (n HtmlNode) getNodesFromChilds(label string)(results []*HtmlNode){
	// 检索自身子节点
	childs := n.GetChilds()
	for _, child := range childs{
		if child.Node.Data == label{
			results = append(results, child)
		}
	}
	return
}
// 在子节点中根据属性获取首个节点
func (n HtmlNode) getNodeByAttrFromChilds(label string, attr string, value string)*HtmlNode {
	// 检索自身子节点
	childs := n.GetChilds()
	for _, child := range childs{
		if child.Node.Data == label{
			attrs := child.Node.Attr
			for _, attr_node := range attrs{
				if attr_node.Key == attr && attr_node.Val == value{
					return child
				}
			}
		}
	}
	return nil
}
// 在子节点中根据属性获取多个节点
func (n HtmlNode) getNodesByAttrFromChilds(label string, attr string, value string)(results []*HtmlNode){
	// 检索自身子节点
	childs := n.GetChilds()
	for _, child := range childs{
		if child.Node.Data == label{
			attrs := child.Node.Attr
			for _, attr_node := range attrs{
				if attr_node.Key == attr && attr_node.Val == value{
					results = append(results, child)
				}
			}
		}
	}
	return
}

// 获取所有子节点
func (n HtmlNode) GetChilds()(nodes []*HtmlNode){
	node := n.Node.FirstChild
	for node != nil{
		nodes = append(nodes, &HtmlNode{node})
		node = node.NextSibling
	}
	return
}
// 获取下一个非文本节点
func (n *HtmlNode) GetNextNotTextCode()*HtmlNode {
	next := n.Node.NextSibling
	for next != nil{
		if next.Type != 1{
			break
		}
		next = next.NextSibling
	}
	return &HtmlNode{next}
}
// 获取Xpath
func (n HtmlNode) Xpath(xpath string)*HtmlNode {
	this_xpath, childs_xpath := n.splitXpathString(xpath)
	// 处理自身
	var node *HtmlNode
	if this_xpath[0:2] == "//"{
		node = n.procXpathString(this_xpath[2:], false)
	}else{
		node = n.procXpathString(this_xpath[1:], true)
	}
	// 处理子节点
	if node != nil && childs_xpath != ""{
		node = node.Xpath(childs_xpath)
	}
	return node
}
// 获取Xpaths
func (n HtmlNode) Xpaths(xpath string)[]*HtmlNode {
	this_xpath, childs_xpath := n.splitXpathString(xpath)
	// 处理自身
	var nodes []*HtmlNode
	if this_xpath[0:2] == "//"{
		nodes = n.procMultiXpathString(this_xpath[2:], false)
	}else{
		nodes = n.procMultiXpathString(this_xpath[1:], true)
	}
	// 处理子节点
	var results []*HtmlNode
	if len(nodes) > 0 && childs_xpath != ""{
		for _, node := range nodes{
			child_nodes := node.Xpaths(childs_xpath)
			for _, child_node := range child_nodes{
				results = append(results, child_node)
			}
		}
		nodes = results
	}
	return nodes
}
// 获取首个节点
func (n *HtmlNode) GetNode(label string)*HtmlNode {
	// 检索自身
	if n.Node.Data == label{
		return n
	}
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		node := child.GetNode(label)
		if node != nil{
			return node
		}
	}
	return nil
}
// 获取多个节点
func (n *HtmlNode) GetNodes(label string)(results []*HtmlNode){
	// 检索自身
	if n.Node.Data == label{
		results = append(results, n)
	}
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		nodes := child.GetNodes(label)
		if len(nodes) > 0{
			for _, node := range nodes{
				results = append(results, node)
			}
		}
	}
	return
}
// 根据属性获取首个节点
func (n *HtmlNode) GetNodeByAttr(label string, attr string, value string)*HtmlNode {
	// 检索自身
	if n.Node.Data == label{
		attrs := n.Node.Attr
		for _, attr_node := range attrs{
			if attr_node.Key == attr && attr_node.Val == value{
				return n
			}
		}
	}
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		node := child.GetNodeByAttr(label, attr, value)
		if node != nil{
			return node
		}
	}
	return nil
}
// 根据属性获取多个节点
func (n *HtmlNode) GetNodesByAttr(label string, attr string, value string)(results []*HtmlNode){
	// 检索自身
	if n.Node.Data == label{
		attrs := n.Node.Attr
		for _, attr_node := range attrs{
			if attr_node.Key == attr && attr_node.Val == value{
				results = append(results, n)
			}
		}
	}
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		nodes := child.GetNodesByAttr(label, attr, value)
		if len(nodes) > 0{
			for _, node := range nodes{
				results = append(results, node)
			}
		}
	}
	return
}
// 根据text获取节点
func (n *HtmlNode) GetNodeByText(value string)*HtmlNode {
	// 检索自身
	for _, text := range n.Texts(){
		if text == value{
			return n
		}
	}
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		node := child.GetNodeByText(value)
		if node != nil{
			return node
		}
	}
	return nil
}
// 根据text获取多个节点
func (n *HtmlNode) GetNodesByText(value string)(results []*HtmlNode){
	// 检索自身
	for _, text := range n.Texts(){
		if text == value{
			results = append(results, n)
		}
	}
	// 检索子节点
	// 检索子节点
	childs := n.GetChilds()
	for _, child := range childs{
		nodes := child.GetNodesByText(value)
		for _, node := range nodes{
			results = append(results, node)
		}
	}
	return
}
// 获取属性值
func (n HtmlNode) GetAttrValue(attr string)string{
	for _, a := range n.Node.Attr{
		if a.Key == attr{
			return a.Val
		}
	}
	return ""
}
// 获取第一个text
func (n HtmlNode) Text()string{
	nodes := n.GetChilds()
	for _, node := range nodes{
		if node.Node.Type == 1{
			return node.Node.Data
		}
	}
	return ""
}
// 获取texts
func (n HtmlNode) Texts()(results []string){
	nodes := n.GetChilds()
	for _, node := range nodes{
		if node.Node.Type == 1{
			results = append(results, node.Node.Data)
		}
	}
	return
}