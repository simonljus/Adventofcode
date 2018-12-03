require 'Set'
class Solver
    attr_accessor :fine
    def initialize(fabric,unclaimed,overlap ="x")
        @fabric=fabric
        @fine=Set.new()
        @overlap =overlap
        @unclaimed =unclaimed
    end
    def addClaims(lines)
        lines.each{|id,x_start,y_start,width,height|
            fine.add(id.to_s)
            height.times{|y|
                width.times{|x|
                    mark= @fabric[y_start+y][x_start+x]
                    if  mark ==@unclaimed
                        @fabric[y_start+y][x_start+x]=id.to_s
                    else
                        fine.delete(id.to_s)
                        fine.delete(mark)
                        @fabric[y_start+y][x_start+x] =@overlap
                    end
                }
            }
        }
    end
    def countMatches()
        area =0
        @fabric.each{|row| area+=row.count{ |cell|cell == @overlap } }
        return area
    end
end
def solve(height=1000,width=1000,unclaimed=".",filename="input.txt")
    fabric = Array.new(height){Array.new(width){unclaimed}}
    solver = Solver.new(fabric,unclaimed)
    fileObj = File.new(filename, "r")
    lines = fileObj.readlines.map{|line| line.split(/[\D]+/)[1..-1].map{
            |v| v.to_i
        }
    }
    puts lines.length
    solver.addClaims(lines)
    area = solver.countMatches()
    puts "Problem 1 solved: #{area} square inches are claimed"
    puts "Problem 2 solved: #{solver.fine.to_a.join("")} is not overlapped"
end
if __FILE__ == $0
    solve()
end
