package xmas

import "testing"

func TestMessage_Crack(t *testing.T) {
	t.Run("it should find the first invalid number ex1", func(t *testing.T) {
		packets := make([]int, 28)
		for i := 0; i < 25; i++ {
			packets[i] = i + 1
		}
		packets[25] = 26
		packets[26] = 49
		packets[27] = 100
		m := &Message{
			packets:        packets,
			preambleLength: 25,
		}
		want := 100
		if got := m.Crack(); got != want {
			t.Errorf("Crack() = %v, want %v", got, want)
		}
	})
	t.Run("it should find the first invalid number ex2", func(t *testing.T) {
		packets := []int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}
		m := &Message{
			packets:        packets,
			preambleLength: 5,
		}
		want := 127
		if got := m.Crack(); got != want {
			t.Errorf("Crack() = %v, want %v", got, want)
		}
	})
}

func BenchmarkMessage_Crack(b *testing.B) {
	nums := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	preambleLength := 5
	message := NewMessage(nums, preambleLength)
	for n := 0; n < b.N; n++ {
		message.Crack()
	}
}

func TestMessage_Exploit(t *testing.T) {
	type fields struct {
		packets        []int
		preambleLength int
	}
	type args struct {
		cracked int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "it should find the sum of the min and max contigous parts",
			fields: fields{
				packets: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
				preambleLength: 0,
			},
			args: args{cracked: 127},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				packets:        tt.fields.packets,
				preambleLength: tt.fields.preambleLength,
			}
			if got := m.Exploit(tt.args.cracked); got != tt.want {
				t.Errorf("Exploit() = %v, want %v", got, tt.want)
			}
		})
	}
}
