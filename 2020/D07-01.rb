require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @rules = IO.read('./D07-input').split("\n").map { |r| Rule.new(r) }
  end

  def run
    @rules.count { |r| r.contains_color?('shiny gold') }
    find_all_parents('shiny gold', @rules).size
  end

  private

  def find_all_parents(color, rules)
    colors = rules.select { |r| r.contains_color?(color) }.map(&:name)
    colors << colors.map { |c| find_all_parents(c, rules) }
    colors.flatten.uniq
  end
end

class Rule
  def initialize(data)
    @raw_data = data
  end

  def name
    @name ||= @raw_data.split(' bags contain ').first
  end

  def contents
    @contents ||= @raw_data.split(' bags contain ').last.split(', ')
  end

  def contains_color?(color)
    contents.any? { |c| c.match?(color) }
  end
end

Measure.new('D07-01', Calculation.new).run
