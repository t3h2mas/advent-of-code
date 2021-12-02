class Sonar
  def self.depth_increases(measurements)
    last_measurement = nil

    increases = 0

    measurements.each do |measurement|
      if last_measurement.nil?
        last_measurement = measurement
        next
      end

      increases += 1 if measurement > last_measurement

      last_measurement = measurement
    end

    increases
  end

  def self.depth_increases_over_window(measurements)
    last_window_measurement = nil
    increases = 0

    for i in 0..measurements.length do
      break if i + 2 >= measurements.length

      window_measurement = measurements[i] + measurements[i + 1] + measurements[i + 2]

      if last_window_measurement.nil?
        last_window_measurement = window_measurement
        next
      end


      increases += 1 if window_measurement > last_window_measurement

      last_window_measurement = window_measurement
    end

    increases
  end
end
