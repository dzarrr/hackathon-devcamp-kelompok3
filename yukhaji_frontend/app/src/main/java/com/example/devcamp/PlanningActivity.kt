package com.example.devcamp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.os.DropBoxManager
import android.util.EventLogTags
import android.util.Half.toFloat
import android.util.Log
import com.android.volley.Request
import com.android.volley.Response
import com.android.volley.toolbox.JsonObjectRequest
import com.android.volley.toolbox.Volley
import com.github.mikephil.charting.charts.LineChart
import com.github.mikephil.charting.components.AxisBase
import com.github.mikephil.charting.components.Description
import com.github.mikephil.charting.components.XAxis
import com.github.mikephil.charting.data.BarData
import com.github.mikephil.charting.data.Entry
import com.github.mikephil.charting.data.LineData
import com.github.mikephil.charting.data.LineDataSet
import com.github.mikephil.charting.utils.ColorTemplate
import com.github.mikephil.charting.components.YAxis
import com.github.mikephil.charting.formatter.IAxisValueFormatter
import kotlinx.android.synthetic.main.activity_planning.*
import org.json.JSONObject

class PlanningActivity : AppCompatActivity() {

    private lateinit var linechart:LineChart
    private var balance = 0
    private var target = 0
    private var timeStart = ""
    private var timeEnd = ""

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_planning)
        linechart = findViewById(R.id.chart) as LineChart



        var a = getPlanningActivity()
        Log.i("HELLO", a.toString())
    }

    private fun createGraphic(balance: Int, target: Int) {
        val xValues: Array<String> = arrayOf("start", "current")
        val yValues: Array<Float> = arrayOf(0f, balance/target.toFloat() * 100)

        val entries = ArrayList<Entry>()
        for ((index, value) in yValues.withIndex()) {
            entries.add(Entry(index.toFloat(), value))
        }

        val labels = ArrayList<String>()
        for (label in xValues) {
            labels.add(label)
        }

        val lineDataSet = LineDataSet(entries, "Progress Tabungan")
        lineDataSet.setColors(R.color.colorPrimary)
        lineDataSet.setDrawFilled(true)


//      Background -> Ensuring the x axis to not be messed up

        val backgroundYValues: Array<Float> = arrayOf(0f, balance/target.toFloat()*100, 100f)

        val backgroundEntries = ArrayList<Entry>()
        for ((index, value) in backgroundYValues.withIndex()) {
            backgroundEntries.add(Entry(index.toFloat(), value ))
        }


        val backgroundLineDataSet = LineDataSet(backgroundEntries, "")
        backgroundLineDataSet.setColors(R.color.colorPrimary)
        backgroundLineDataSet.setDrawFilled(false)
        backgroundLineDataSet.enableDashedLine(0f, 1f, 0f)
        backgroundLineDataSet.setDrawCircles(false)
        backgroundLineDataSet.setDrawValues(false)

        var lineData = LineData(lineDataSet)
        lineData.addDataSet(backgroundLineDataSet)

        var xAxis: XAxis = linechart.getXAxis()

        xAxis.setDrawAxisLine(false)
        xAxis.setDrawGridLines(false)
        xAxis.setLabelCount(0)
        xAxis.setPosition(XAxis.XAxisPosition.BOTTOM)
        xAxis.setValueFormatter(ChartTimeFormatter())



        val yAxis = linechart.getAxisLeft()
        val yAxisRight = linechart.getAxisRight()

        yAxis.setLabelCount(3, true)
        yAxis.setDrawGridLines(false)
        yAxis.setDrawZeroLine(true)
        yAxis.setDrawAxisLine(false)
        yAxis.setAxisMinimum(0f)
        yAxis.setAxisMaximum(100f)
        yAxis.setDrawLabels(false)
        yAxisRight.setLabelCount(3, true)
        yAxisRight.setDrawGridLines(false)
        yAxisRight.setDrawZeroLine(true)
        yAxisRight.setDrawAxisLine(true)
        yAxisRight.setAxisMinimum(0f)
        yAxisRight.setAxisMaximum(100f)

        linechart.setTouchEnabled(false)
        linechart.setDragEnabled(false)
        linechart.setScaleEnabled(false)
        linechart.setDrawGridBackground(false)
        linechart.setPinchZoom(false)
        val description = Description()
        description.text = ""
        linechart.description = description
        linechart.legend.setEnabled(false)

        linechart.setData(lineData)
        linechart.invalidate()
    }

    private fun getPlanningActivity(){
        val queue = Volley.newRequestQueue(this)
        val API_url = "http://10.50.218.18:3000/saving/$user_id"
        val jsonObjectRequest = JsonObjectRequest(
            Request.Method.GET, API_url, null,
            Response.Listener { response ->
                var strResp = response.toString()
                val jsonObj: JSONObject = JSONObject(strResp)
                balance = jsonObj.getString("balance").toInt()
                target = jsonObj.getString("target").toInt()
                timeStart = jsonObj.getString("start_date")
                timeEnd = jsonObj.getString("end_date")
                createGraphic(balance, target)
                Log.d("balance", balance.toString())
                tvRencanaHaji.text = timeEnd.slice(0..3)
                var tahunAkhir = tvRencanaHaji.text.toString().toInt()
                var sisaWaktu = tahunAkhir - 2019
                var bulanan = (target - balance) / sisaWaktu * 12
                tvBiayaBulanan.text = "Rp ${bulanan.toString()} / Bulan"
                Log.d("target", target.toString())
                Log.d("time_start", timeStart)
                Log.d("time_end", timeEnd)
            },
            Response.ErrorListener { error ->
                Log.d("MyTag", error.message)
            }
        )
        queue.add(jsonObjectRequest)
    }
}



class ChartTimeFormatter : IAxisValueFormatter {

    override fun getFormattedValue(value: Float, axis: AxisBase?): String = when (value) {
        0f -> "Start"
        1f -> "Current"
        2f -> "End"
        else -> ""
    }
}
