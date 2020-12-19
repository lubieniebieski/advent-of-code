require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @boarding_passes = IO.read('./D05-input').split("\n")
    @boarding_passes = @boarding_passes.map { |b| BoardingPass.new(b) }
  end

  def run
    @boarding_passes.map(&:seat_id).max
  end
end

class BoardingPass
  def initialize(string)
    @raw_data = string
  end

  def row
    to_i(@raw_data[0..6])
  end

  def column
    to_i(@raw_data[7..9])
  end

  def seat_id
    row * 8 + column
  end

  def to_s
    "#{@raw_data}: row #{row}, column #{column}, seat ID #{seat_id}."
  end

  private

  def to_i(string)
    string.tr('F', '0').tr('B', '1').tr('R', '1').tr('L', '0').to_i(2)
  end
end

Measure.new('D05-01', Calculation.new).run
