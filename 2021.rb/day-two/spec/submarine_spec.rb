require 'submarine'

RSpec.describe Submarine do
  describe 'using DirectionalNavigation' do
    let(:submarine) { Submarine.new(DirectionalNavigation.new) }

    context 'when moving horizontally' do
      it 'moves forward' do
        submarine.move(:forward, 1)
        expect(submarine.horizontal_position).to eq(1)
      end

      it 'moves backward' do
        submarine.move(:backward, 1)
        expect(submarine.horizontal_position).to eq(-1)
      end
    end

    context 'when moving vertically' do
      it 'moves up' do
        submarine.move(:up, 1)
        expect(submarine.depth).to eq(-1)
      end

      it 'moves down' do
        submarine.move(:down, 1)
        expect(submarine.depth).to eq(1)
      end
    end

    context 'when moving in all directions' do
      let(:commands) do
        [
          { direction: :forward, count: 5 },
          { direction: :down, count: 5 },
          { direction: :forward, count: 8 },
          { direction: :up, count: 3 },
          { direction: :down, count: 8 },
          { direction: :forward, count: 2 }
        ]
      end

      it 'ends where expected' do
        commands.each do |command|
          submarine.move(command[:direction], command[:count])
        end
        expect(submarine.horizontal_position).to eq(15)
        expect(submarine.depth).to eq(10)
      end
    end
  end
end
