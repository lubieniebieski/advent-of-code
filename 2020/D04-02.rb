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
  KEY_NAMES = %w(byr iyr eyr hgt hcl ecl pid cid)
  OPTIONAL_KEYS = %w(cid)

  attr_reader :data

  def initialize(data)
    raw_data = data.gsub("\n", " ").split(" ")
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
    correct_fields_present? &&
      correct_years? &&
      valid_height? &&
      valid_hair? &&
      valid_eye_color? &&
      valid_pid?
  end

  def invalid?
    !valid?
  end

  private

  def correct_years?
    (1920..2002).include?(@data["byr"].to_i) &&
    (2010..2020).include?(@data["iyr"].to_i) &&
    (2020..2030).include?(@data["eyr"].to_i)
  end

  def valid_height?
    @data["hgt"].match?(/1([5-8]\d|9[0-3])cm|(59|6\d|7[0-6])in/)
  end

  def valid_hair?
    @data["hcl"].match(/#[0-9a-f]{6}/)
  end

  def valid_eye_color?
    @data["ecl"].match(/(amb|blu|brn|gry|grn|hzl|oth)/)
  end

  def valid_pid?
    @data["pid"].match?(/\d{9}/)
  end

  def correct_fields_present?
    shortcut == KEY_NAMES.sort.join ||
      shortcut == (KEY_NAMES - OPTIONAL_KEYS).sort.join
  end
end

Measure.new('D04-02', Calculation.new).run
