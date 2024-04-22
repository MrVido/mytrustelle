"use client";

import Footer from "@/components/v1/layout/Footer";
import Navbar from "@/components/v1/layout/layout2/Navbar.client";
import React, { useState } from "react";
import { Button } from "@/components/v1/ui/button";
import { cn } from "@/lib/utils";
import { Input } from "@/components/ui/Input";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable"
import Navigation from "@/components/navbar/navigation";

export default function Home() {
  const handleMenuButtonClick = () => {
    console.log("Menu button clicked");
  };
  return (
    <div>
      <Navigation />
              <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <div>
             <ResizablePanelGroup
      direction="horizontal"
      className="max-w-md rounded-lg border"
    >
      <ResizablePanel defaultSize={50}>
        <div className="flex h-[200px] items-center justify-center p-6">
          <span className="font-semibold">One</span>
        </div>
      </ResizablePanel>
      <ResizableHandle />
      <ResizablePanel defaultSize={50}>
        <ResizablePanelGroup direction="vertical">
          <ResizablePanel defaultSize={25}>
            <div className="flex h-full items-center justify-center p-6">
              <span className="font-semibold">Two</span>
            </div>
          </ResizablePanel>
          <ResizableHandle />
          <ResizablePanel defaultSize={75}>
            <div className="flex h-full items-center justify-center p-6">
              <span className="font-semibold">Three</span>
            </div>
          </ResizablePanel>
        </ResizablePanelGroup>
      </ResizablePanel>
    </ResizablePanelGroup>
        </div>
      </main>
        <div className="flex content-center justify-center relative gap-4 p-10 m-10">
        <Footer />
      </div>
    </div>
  );
}
