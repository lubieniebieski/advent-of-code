require_relative '../solutions/D11'

describe D11 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D11.txt') }
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
      '5483143223
      2745854711
      5264556173
      6141336146
      6357385478
      4167524645
      2176841721
      6882881134
      4846848554
      5283751526'
    end
    let(:test_input) { test_data.split("\n") }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 1656
      end
      it 'Matches 20 step data' do
        step_data = "3936556452
        5686556806
        4496555690
        4448655580
        4456865570
        5680086577
        7000009896
        0000000344
        6000000364
        4600009543".split("\n")
        expect(subject.compare(step_data, 20)).to eq true
      end

      context 'number of flashes' do
        it { expect(subject.flashes_in_step(1)).to eq 0 }
        it { expect(subject.flashes_in_step(2)).to eq 35 }
        it { expect(subject.flashes_in_step(3)).to eq 45 }
        it { expect(subject.flashes_in_step(4)).to eq 16 }
        it { expect(subject.flashes_in_step(5)).to eq 8 }
        it { expect(subject.flashes_in_step(100)).to eq 13 }
      end

      it 'Matches 60 step data' do
        step_data = "2533334200
        2743334640
        2264333458
        2225333337
        2225333338
        2287833333
        3854573455
        1854458611
        1175447111
        1115446111".split("\n")
        expect(subject.compare(step_data, 60)).to eq true
      end

      it 'Matches 100 step data' do
        step_data = "0397666866
        0749766918
        0053976933
        0004297822
        0004229892
        0053222877
        0532222966
        9322228966
        7922286866
        6789998766".split("\n")
        expect(subject.compare(step_data, 100)).to eq true
      end

      it 'Matches 10 step data' do
        step_data = "0481112976
        0031112009
        0041112504
        0081111406
        0099111306
        0093511233
        0442361130
        5532252350
        0532250600
        0032240000".split("\n")
        expect(subject.compare(step_data, 10)).to eq true
      end

      it 'Matches 5 step data' do
        step_data = "4484144000
        2044144000
        2253333493
        1152333274
        1187303285
        1164633233
        1153472231
        6643352233
        2643358322
        2243341322".split("\n")
        expect(subject.compare(step_data, 5)).to eq true
      end

      it 'Matches 193 step data' do
        step_data = "5877777777
        8877777777
        7777777777
        7777777777
        7777777777
        7777777777
        7777777777
        7777777777
        7777777777
        7777777777".split("\n")
        expect(subject.compare(step_data, 193)).to eq true
      end

      it 'Matches 3 step data' do
        step_data = "0050900866
        8500800575
        9900000039
        9700000041
        9935080063
        7712300000
        7911250009
        2211130000
        0421125000
        0021119000".split("\n")
        expect(subject.compare(step_data, 3)).to eq true
      end

      it 'Matches 4 step data' do
        step_data = "2263031977
        0923031697
        0032221150
        0041111163
        0076191174
        0053411122
        0042361120
        5532241122
        1532247211
        1132230211".split("\n")
        expect(subject.compare(step_data, 4)).to eq true
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 195
      end
    end
  end
end
