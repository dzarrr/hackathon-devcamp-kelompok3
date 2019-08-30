package com.example.devcamp

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        btnMulaiNabung.setOnClickListener {
            startNewActivity()
        }
    }

    private fun startNewActivity() {
        val intent = Intent(this, PhotoActivity::class.java)
        startActivity(intent)


//        if (adaTabungan == false) {
//            val intent = Intent(this, PhotoActivity::class.java)
//            startActivity(intent)
//        } else {
//            val intent = Intent(this, TargetActivity::class.java)
//            startActivity(intent)
//        }

    }
}
