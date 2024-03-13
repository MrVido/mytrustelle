'use client'

import Footer from "@/components/layout/Footer";
import Navbar from "@/components/layout/layout2/Navbar.client";
import React, { useState } from "react";
import { Button } from "@/components/ui/button"

import Image from "next/image";

export default function Home() {
  const handleMenuButtonClick = () => {
    console.log("Menu button clicked");
  };
  return (
    <div>
   <Navbar onMenuButtonClick={handleMenuButtonClick} />
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <p className="fixed left-0 top-0 flex w-full justify-center border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-6 pt-8 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:static lg:w-auto  lg:rounded-xl lg:border lg:bg-gray-200 lg:p-4 lg:dark:bg-zinc-800/30">
          Get started by editing&nbsp;
          <code className="font-mono font-bold">src/app/page.tsx</code>
          
        </p>
        <Button>Click me</Button>
      </div>
      <Footer />
    </main>
    </div>
  );
}
