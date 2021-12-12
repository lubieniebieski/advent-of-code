require_relative '../solutions/D12'

describe D12 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D12.txt') }
    let(:input) { File.read(input_path).split("\n") }
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
      results = described_class.run(input)
      warn "\t#{results}"
      expect(results).not_to eq nil
    end
  end

  context 'Provided examples' do
    let(:test_data) do
      'start-A
      start-b
      A-c
      A-b
      b-d
      A-end
      b-end'
    end
    let(:test_input) { test_data.split("\n") }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 10
      end
      it 'Passes the 2nd example' do
        test_input = "dc-end
        HN-start
        start-kj
        dc-start
        dc-HN
        LN-dc
        HN-end
        kj-sa
        kj-HN
        kj-dc"
        expect(described_class.new(test_input.split("\n")).part1).to eq 19
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 36
      end
    end
  end
end
