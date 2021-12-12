require 'pry'

class D12
  def initialize(data)
    @nodes = {}
    data.map(&:strip).map do |pair|
      a, b = pair.split('-')
      @nodes[a] ||= []
      @nodes[a] << b
      @nodes[b] ||= []
      @nodes[b] << a
    end
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    @paths = []
    traverse_node('start', ['start'])
    @paths.size
  end

  def part2
    @paths = []
    small_caves = @nodes.keys.select { |node| node.downcase == node } - %w[start end]
    small_caves.each do |cave|
      traverse_node('start', ['start'], cave)
    end
    @paths.uniq.size
  end

  def traverse_node(start_node, visited_nodes = [], duplicated_cave = nil)
    @nodes[start_node].each do |hop_name|
      if hop_name == 'end'
        @paths << (visited_nodes + [hop_name])
        next
      end

      if visited_nodes.include?(hop_name) && hop_name.downcase == hop_name
        if hop_name == duplicated_cave
          next if visited_nodes.count(hop_name) > 1
        else
          next
        end
      end

      traverse_node(hop_name, (visited_nodes + [hop_name]), duplicated_cave)
    end
  end
end
