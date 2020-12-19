require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @passports = IO.read('./D04-input').split("\n\n")
    @passports = @passports.map { |p| Passport.new(p) }
  end

  def run
    @passports.count(&:valid?)
  end
end

class Passport
  KEY_NAMES = %w(byr iyr eyr hgt hcl ecl pid cid).freeze
  OPTIONAL_KEYS = %w(cid).freeze

  attr_reader :data

  def initialize(data)
    raw_data = data.tr("\n", ' ').split(' ')
    @data = {}
    raw_data.each do |rd|
      key, value = rd.split(':')
      @data[key] = value
    end
  end

  def shortcut
    @data.keys.sort.join
  end

  def valid?
    shortcut == KEY_NAMES.sort.join ||
      shortcut == (KEY_NAMES - OPTIONAL_KEYS).sort.join
  end

  def invalid?
    !valid?
  end
end

Measure.new('D04-01', Calculation.new).run
