# typed: strict
require 'set'

require 'sorbet-runtime'

class Caves
  extend T::Sig

  sig { params(input: String).returns(Caves) }
  def self.parse(input)
    g = new

    input.split("\n").map { |l| l.split('-') }.each do |connection|
      start, destination = connection
      g.connect(start, destination)
    end

    g
  end

  sig { returns(T::Hash[String, Cave]) }
  attr_reader :nodes

  sig { void }
  def initialize
    @nodes = T.let({}, T::Hash[String, Cave])
  end

  sig { params(head: String, tail: String).void }
  def connect(head, tail)
    head_node = add(head)
    tail_node = add(tail)

    head_node.connect_to(tail_node)
    tail_node.connect_to(head_node)
  end

  sig { params(value: String).returns(Cave) }
  def add(value)
    existing = @nodes[value]

    return existing unless existing.nil?

    node = Cave.new(value)
    @nodes[value] = node
    node
  end
end

class Cave
  extend T::Sig

  sig { returns(String) }
  attr_reader :value

  sig { returns(T::Hash[String, Cave]) }
  attr_reader :edges

  sig { params(value: String).void }
  def initialize(value)
    @value = value
    @edges = T.let({}, T::Hash[String, Cave])
  end

  sig { params(node: Cave).void }
  def connect_to(node)
    @edges[node.value] = node
  end

  sig { returns(T::Boolean) }
  def end?
    @value == 'end'
  end

  sig { returns(T::Boolean) }
  def small?
    !end? && @value == @value.downcase
  end

  sig { returns(T::Boolean) }
  def big?
    !end? && @value == @value.upcase
  end

  sig { returns(String) }
  def to_s
    "Cave<'#{@value}' edges: #{@edges.count}>"
  end
end

class Explorer
  extend T::Sig

  sig { params(caves: Caves).returns(Integer) }
  def self.paths(caves)
    starting_cave = caves.nodes['start']

    raise 'Failed to find start' if starting_cave.nil?

    starting_cave.edges.values.map do |c|
      visited = T.let(Set.new(['start']), T::Set[String])
      paths_from(c, visited)
    end.sum
  end

  sig { params(cave: Cave, visited: T::Set[String]).returns(Integer) }
  def self.paths_from(cave, visited)
    return 1 if cave.end?
    return 0 if visited.include?(cave.value)

    puts "Visiting cave: #{cave}"

    visited << cave.value if cave.small?

    cave.edges.values.map { |c| paths_from(c, visited) }.sum
  end
end
