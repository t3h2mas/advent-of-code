input = <<~TEXT
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
TEXT

input = File.read('./07_input.txt')

equations = input.split("\n").map do |line| 
  target, test_inputs = line.split(': ')
  { target: target.to_i, test_inputs: test_inputs.split(' ').map(&:to_i) }
end;

def build_evaluations(xs)
  root = Value.new(value: xs.first, edges: [])

  build_evaluation_steps(root, xs[1..])

  root
end

def build_evaluation_steps(parent, xs)
  return if xs.empty?

  value = xs.first

  new_sum = Value.new(value:, edges: [])
  new_mul = Value.new(value:, edges: [])
  new_cat = Value.new(value:, edges: [])

  parent.edges << Edge.new(operator: :+, operand: new_sum)
  parent.edges << Edge.new(operator: :*, operand: new_mul)
  parent.edges << Edge.new(operator: :&, operand: new_cat)

  build_evaluation_steps(new_sum, xs[1..])
  build_evaluation_steps(new_mul, xs[1..])
  build_evaluation_steps(new_cat, xs[1..])
end

Value = Data.define(:value, :edges) do
  def +(other)
    self.value + other.value
  end

  def *(other)
    self.value * other.value
  end

  def &(other)
    "#{self.value}#{other.value}".to_i
  end
end

Edge = Data.define(:operator, :operand)

tree = build_evaluations([1,2,3])

def path_total?(node, target, total = nil)
  return false if !total.nil? && total > target
  return total == target if node.edges.empty?

  total = node.value if total.nil?

  total_sum = Value.new(value: total, edges: []).send(node.edges[0].operator, node.edges[0].operand)
  total_mul = Value.new(value: total, edges: []).send(node.edges[1].operator, node.edges[1].operand)
  total_cat = Value.new(value: total, edges: []).send(node.edges[2].operator, node.edges[2].operand)

  path_total?(node.edges[0].operand, target, total_sum) || path_total?(node.edges[1].operand, target, total_mul) || path_total?(node.edges[2].operand, target, total_cat)
end

equations.filter do |equation|
  tree = build_evaluations(equation[:test_inputs])
  path_total?(tree, equation[:target])
end.sum { _1[:target] }
