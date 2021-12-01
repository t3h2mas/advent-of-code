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
end
