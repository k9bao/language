package com.example.mywatch;

import android.app.Activity;
import android.content.pm.ActivityInfo;
import android.graphics.Canvas;
import android.graphics.Color;
import android.graphics.Paint;
import android.os.Bundle;
import android.view.SurfaceHolder;
import android.view.SurfaceView;
import android.view.WindowManager;

import java.text.SimpleDateFormat;
import java.util.Date;

public class MainActivity extends Activity {

    SurfaceView mSurfaceView;
    SurfaceHolder mHolder;
    Thread mThread;
    Paint mPaint;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        getWindow().addFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN);
        setRequestedOrientation(ActivityInfo.SCREEN_ORIENTATION_LANDSCAPE);

        setContentView(R.layout.activity_main);

        mSurfaceView = findViewById(R.id.surfaceView);
        mSurfaceView.getHolder().addCallback(new SurfaceHolder.Callback() {
            @Override
            public void surfaceCreated(SurfaceHolder holder) {
                mHolder =  holder;
            }

            @Override
            public void surfaceChanged(SurfaceHolder holder, int format, int width, int height) {
            }

            @Override
            public void surfaceDestroyed(SurfaceHolder holder) {
            }
        });

        mPaint = new Paint();
        mPaint.setTextSize(240);
        mPaint.setColor(Color.BLACK);
    }

    SimpleDateFormat formatter = new SimpleDateFormat("HH:mm:ss.SSS");

    private void render() {
        while (!Thread.interrupted()) {
            if(mHolder != null) {
                String text = formatter.format(new Date());
                Canvas canvas = mHolder.lockCanvas();
                canvas.drawColor(Color.WHITE);
                canvas.drawText(text, 20, 400, mPaint);
                mHolder.unlockCanvasAndPost(canvas);
            }
            Thread.yield();
        }
    }

    @Override
    protected void onStart() {
        super.onStart();
        mThread = new Thread(MainActivity.this::render);
        mThread.start();
    }

    @Override
    protected void onStop() {
        super.onStop();
        if(mThread != null) {
            mThread.interrupt();
        }
    }
}
