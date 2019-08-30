package com.example.devcamp

import android.content.Intent
import android.graphics.Color
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.view.View
import com.android.volley.Request
import com.android.volley.Response
import com.android.volley.toolbox.JsonObjectRequest
import com.android.volley.toolbox.Volley
import kotlinx.android.synthetic.main.activity_target.*
import org.json.JSONObject
import java.text.SimpleDateFormat
import java.util.*

class TargetActivity : AppCompatActivity() {

    var tahun = 2020
    var tahunSekarang = 2019
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_target)

        val creditScore = 90
        if (creditScore > 80) {
            btnNotif.visibility = View.VISIBLE
        } else {
            btnNotif.visibility = View.GONE
        }


        tbInvestasi.setOnClickListener {
            checkInvestasi()
        }

        tbNonInvestasi.setOnClickListener {
            checkNonInvestasi()
        }

        btnSelesai.setOnClickListener {
            postTargetActivity()
            val intent = Intent(this@TargetActivity, PlanningActivity::class.java)
            startActivity(intent)
        }

        btnKurangTahun.setOnClickListener {
            tahun = tvTahun.text.toString().toInt()
            if(tahun == 2020) {
                tahun = 2020
            } else {
                tahun -= 1
            }
            tvTahun.text = tahun.toString()
            displayBiaya()
        }

        btnTambahTahun.setOnClickListener {
            tahun = tvTahun.text.toString().toInt()
            tahun += 1
            tvTahun.text = tahun.toString()
            displayBiaya()
        }
    }

    private fun checkInvestasi() {
        if (tbInvestasi.isChecked == true) {
            tbNonInvestasi.isChecked = false
            tbInvestasi.setTextColor(Color.WHITE)
            tbInvestasi.isEnabled = false
            tbNonInvestasi.isEnabled = true
            tbNonInvestasi.setTextColor(Color.BLACK)
        }
    }

    private fun checkNonInvestasi(){
        if(tbNonInvestasi.isChecked == true) {
            tbInvestasi.isChecked = false
            tbNonInvestasi.setTextColor(Color.WHITE)
            tbNonInvestasi.isEnabled = false
            tbInvestasi.isEnabled = true
            tbInvestasi.setTextColor(Color.BLACK)
        }
    }

    private fun displayBiaya() {
        var jum = 35000000L
        var biaya = tahun - tahunSekarang
        biaya *= 12
        jum /= biaya
        tvTotalBiaya.text = jum.toString()
    }

    private fun postTargetActivity(){
        val jsonBodyObj = JSONObject()
        jsonBodyObj.put("user_id", user_id)
        jsonBodyObj.put("balance", 0)
        jsonBodyObj.put("target", 35000000)
        jsonBodyObj.put("start_date", SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss'Z'").format(Date()).toString())
        jsonBodyObj.put("end_date", "$tahun-01-01T00:00:00Z")

        val queue = Volley.newRequestQueue(this)

        val API_url = "http://10.50.218.18:3000/saving"
        val jsonObjectRequest = JsonObjectRequest(
            Request.Method.POST, API_url, jsonBodyObj,
            Response.Listener { response ->
                Log.d("MyTag", "Response: %s".format(response.toString()))
            },
            Response.ErrorListener { error ->
                Log.d("MyTag", error.message)
            }
        )

        queue.add(jsonObjectRequest)
    }

}
