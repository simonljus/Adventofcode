import java.io.File
import java.io.InputStream


fun fileToString(filename: String): String{
    val inputStream: InputStream = File(filename).inputStream()
 
	val inputString = inputStream.bufferedReader().use { it.readText() }
	return inputString

}

fun sumOfMatchingDistance(cipher: String, jump: Int): Int{
    val n = cipher.length
    var sum =0 
    val zero = '0'.toInt()
    for ((index,value) in cipher.withIndex()){
        when (value){
            cipher[(index+jump) %n] -> sum +=  value.toInt() -zero
        }
    }
    return sum
}

fun main(args: Array<String>) {
    val cipher = fileToString("input.txt")
    val sum = sumOfMatchingDistance(cipher,1)
    val part2 = sumOfMatchingDistance(cipher,cipher.length/2)
    println(sum)
    println(part2)
}