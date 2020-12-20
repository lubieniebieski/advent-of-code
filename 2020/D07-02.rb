require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @rules = IO.read('./D07-input').split("\n").map { |r| Rule.new(r) }
  end

  def run
    bags_inside('shiny gold', @rules) - 1
  end

  private

  def find_all_parents(color, rules)
    colors = rules.select { |r| r.contains_color?(color) }.map(&:name)
    colors << colors.map { |c| find_all_parents(c, rules) }
    colors.flatten.uniq
  end

  def bags_inside(color, rules)
    rule = rules.find { |r| r.name == color }
    return 1 if rule.contents.empty?
    rule.contents.map do |c|
      bags_inside(c.color, rules) * c.quantity
    end.sum + 1
  end
end

class Rule
  attr_reader :contents, :name
  def initialize(data)
    @contents = setup_contents(data)
    @name = data.split(' bags contain ').first
  end

  def name
    @name ||= @raw_data.split(' bags contain ').first
  end

  def contains_color?(color)
    contents.any? { |c| c.color.match?(color) }
  end

  private

  def setup_contents(data)
    c = data.split(' bags contain ').last.split(', ')
    return [] if c.first.match?('no other bags')
    c.map { |co| Content.new(co) }
  end
end

class Content
  attr_reader :color, :quantity

  def initialize(string)
    array = string.split
    array.pop
    @quantity = array.shift.to_i
    @color = array.join(' ')
  end
end

Measure.new('D07-02', Calculation.new).run
