import React from "react";
import Link from "next/link";

export default function LearnMore() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center px-4 py-12">
      <div className="max-w-4xl mx-auto text-center">
        {/*Header*/}
        <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-6">
          Thera<span className="text-blue-600">Map</span>
        </h1>
        
        {/*Description*/}
        <p className="text-xl md:text-2xl text-gray-700 mb-10 max-w-3xl mx-auto leading-relaxed">
          Discover how TheraMap can help you track your emotions, understand your thought patterns, 
          and develop healthier coping strategies through AI-powered insights.
        </p>
        
        {/*Features*/}
        <div className="grid md:grid-cols-3 gap-6 mb-12">
          <div className="bg-white p-6 rounded-xl shadow-md">
            <h3 className="text-xl font-semibold mb-3 text-blue-600">Emotion Tracking</h3>
            <p className="text-gray-600">Log and analyze your daily emotions to identify patterns and triggers.</p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-md">
            <h3 className="text-xl font-semibold mb-3 text-blue-600">Thought Mapping</h3>
            <p className="text-gray-600">Visualize your thought processes and identify negative patterns.</p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-md">
            <h3 className="text-xl font-semibold mb-3 text-blue-600">AI Insights</h3>
            <p className="text-gray-600">Get personalized recommendations based on your unique patterns.</p>
          </div>
        </div>
        
        {/*CTA Buttons*/}
        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Link
            href="/get-started"
            className="px-8 py-4 bg-blue-600 text-white text-lg font-semibold rounded-lg shadow-md hover:bg-blue-700 transition-colors"
          >
            Get Started
          </Link>
          <Link
            href="/"
            className="px-8 py-4 bg-white text-blue-600 text-lg font-semibold rounded-lg shadow-md border border-blue-200 hover:bg-blue-50 transition-colors"
          >
            Back to Home
          </Link>
        </div>
      </div>
    </div>
  );
}