require 'benchmark'
require 'date'

def split(s)
  p = s.split('/').reject(&:empty?).slice(-2..-1)
  return "" if not p
  return p.join('/')
end

def split_reverse(s, n)
  p = s.reverse().split('/').reject(&:empty?)
  return "" if p.length < n
  return p[0..n-1].join('/').reverse()
end

def regexp(s)
  m = %r{[^/]+/+[^/]+\.\w+$}.match(s)
  return "" if not m
  return m[0]
end

def regexp_reverse(s)
  m = %r{^\w+\.[^/]+/+[^/]+}.match(s.reverse())
  return "" if not m
  return m[0].reverse()
end

path = '/regular/debugger/missing/gateways/reaching/major/nurse.js'

iterations = 200_000

Benchmark.bmbm do |bm|
  bm.report('split') do
    iterations.times do
      split(path)
    end
  end
  bm.report('split_reverse') do
    iterations.times do
      split_reverse(path, 2)
    end
  end
  bm.report('regexp') do
    iterations.times do
      regexp(path)
    end
  end
  bm.report('regexp_reverse') do
    iterations.times do
      regexp_reverse(path)
    end
  end
end
