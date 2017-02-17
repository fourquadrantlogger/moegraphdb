import java.io.PrintWriter

import scala.io.Source


/**
  * Created by paidian on 17-2-15.
  */
object MapCsv extends App {
  val file = Source.fromFile("/home/paidian/下载/1000000_10_fans_1000.map")
  val fstr=file.mkString
  var o=fstr.replace(':',',')
  o= o.replace(",\"","\n")
  o= o.replace("\"","")
  val out = new PrintWriter("outputdata.csv")
  out.write(o)
  out.close()
}