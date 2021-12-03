require_relative '../solutions/D09'

describe D09 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D09.txt') }
    let(:input) { IO.read(input_path).split("\n") }
    let(:subject) { described_class.new(input) }
    let(:part1) { subject.part1 }
    let(:part2) { subject.part2 }

    it 'Solves Part 1' do
      expect { part1 }.not_to raise_exception
    end

    it 'Solves Part 2' do
      expect { part2 }.not_to raise_exception
    end

    it 'Works with real data' do
      described_class.run(input)
    end
  end

  context 'Provided examples' do
    let(:test_data) do
      '35
      20
      15
      25
      47
      40
      62
      55
      65
      95
      102
      117
      150
      182
      127
      219
      299
      277
      309
      576'
    end
    let(:test_input) { test_data.split("\n") }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1(5)).to eq 127
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2(5)).to eq 62
      end
    end
  end
end
