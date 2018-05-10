require 'benchmark'
require 'date'

def time_to_millis(t)
  (t.to_f * 1000).to_i
end

def millis_to_time(n)
  Time.at(n / 1000)
end

def rfc3339_to_time(s)
  DateTime::parse(s).to_time
end

iterations = 100_000

$now = Time.now()
$millis = time_to_millis($now)
$rfc3339 = $now.to_datetime().rfc3339()

Benchmark.bmbm do |bm|
  bm.report('millis_to_time') do
    iterations.times do
      millis_to_time($millis)
    end
  end
  bm.report('rfc3339_to_time') do
    iterations.times do
      rfc3339_to_time($rfc3339)
    end
  end
end
