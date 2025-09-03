import React from "react";
import Link from "next/link";

export default function Welcome() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-9 text-center">
      <h1 className="text-5xl font-bold mb-6">
        Welcome to Thera<span className="text-primary-500">Map</span>

      </h1>

      <p className="max-w-2xl text-lg text-gray-600 dark:text-gray-300 mb-8">
        TheraMap is an intelligent system that helps you track your emotions,
        map your thoughts, and transform negative thinking patterns into
        positive, actionable steps—giving you mental clarity and inner peace.
      </p>

      <div className="flex gap-4">
        <button className="px-6 py-3 bg-[#154D71] text-white rounded-lg hover:bg-[#0e3a57] transition-colors">
          Get Started
        </button>

        <Link
          href="/learn-more"
          className="px-6 py-3 bg-[#468A9A] text-white rounded-xl shadow-md hover:bg-[#3a6f7c] transition"
        >
          Learn More
        </Link>
      </div>
    </div>
  );
}