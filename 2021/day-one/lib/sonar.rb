class Sonar
  def self.depth_increases_over_window_two(measurements, window_size)
    increases = 0
    last_measurement = nil

    for i in 0..measurements.length do
      measurement = 0

      break if i + window_size > measurements.length

      for w in 0...window_size do
        measurement += measurements[i + w]
      end

      if last_measurement.nil?
        last_measurement = measurement
        next
      end

      increases += 1 if measurement > last_measurement
      last_measurement = measurement
    end

    increases
  end
end
