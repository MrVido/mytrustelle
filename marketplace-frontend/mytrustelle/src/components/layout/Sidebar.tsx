import React, { useRef } from "react";
import classNames from "classnames";
import Link from "next/link";
import Image from "next/image";
import { defaultNavItems } from "./defaultNavItems";
import { useOnClickOutside } from "usehooks-ts";
import CarGirlsLogo from "/public/CarGirls-WebLogo.png";
import { signIn, signOut, useSession } from "next-auth/react";



// define a NavItem prop
export type NavItem = {
  label: string;
  href: string;
  icon: React.ReactNode;
};
// add NavItem prop to component prop
type Props = {
  open: boolean;
  navItems?: NavItem[];
  setOpen(open: boolean): void;
};
const Sidebar = ({ open, navItems = defaultNavItems, setOpen }: Props) => {
  const ref = useRef<HTMLDivElement>(null);
  useOnClickOutside(ref, (e) => {
    setOpen(false);
  });
  return (
    <div
      className={classNames({

        "flex flex-col justify-between": true, // layout
        "bg-indigo-700 text-zinc-50": true, // colors
        "lg:w-[160px] lg:sticky lg:top-16 lg:z-0 top-0 z-20 fixed": true, // positioning
        "lg:h-[calc(100vh_-_64px)] h-full w-[120px]": true, // for height and width
        "transition-transform .3s ease-in-out lg:-translate-x-0": true,
        "md:w-[160px] md:sticky md:top-16 md:z-0 top-0 z-20 fixed": true, // positioning
        "md:h-[calc(100vh_-_64px)] h-full w-[120px]": true, // for height and width
        "transition-transform .3s ease-in-out md:-translate-x-0": true,
        "sm:h-[calc(100vh_-_24px)] h-full w-[120px]": true, // for height and width
        "sm:w-[160px] sm:sticky sm:top-10 sm:z-0 top-0 z-20 overflow": true,
        "transition-transform .3s ease-in-out sm:-translate-x-0": true,
        "xs:h-[calc(100vh_-_64px)] h-full w-[10px]": true, // for height and width
        "xs:w-full xs:sticky xs:top-10 xs:z-0 top-0 z-20 overflow": true,
        "transition-transform .3s ease-in-out xs:-translate-x-0": true,
        "-translate-x-4 ": !open, //hide sidebar to the left when closed
      })}
      ref={ref}
    >
      <nav className="md:sticky top-0 md:top-16 sm:sticky top:-2">
        {/* nav items */}
        <ul className="py-2 flex flex-col gap-2">
          {navItems.map((item, index) => {
            return (
              <Link key={index} href={item.href}>
                <li
                  className={classNames({
                    "text-indigo-100 hover:bg-indigo-900": true, //colors
                    "flex gap-4 items-center": true, //layout
                    "transition-colors duration-300": true, //animation
                    "rounded-md p-2 mx-2": true, //self style
                  })}
                >
                  {item.icon} {item.label}
                </li>
              </Link>
            );
          })}
        </ul>
      </nav>
      {/* Chatbot  */}
      {/* account  */}
      <div className="border-t border-t-indigo-800 p-4">
        <div className="flex gap-4 items-center">
          <Image
            src={CarGirlsLogo}
            height={36}
            width={36}
            alt="profile image"
            className="rounded-full"
          />
          <div className="flex flex-col ">
            <span className="text-indigo-50 my-0"><AuthShowcase /></span>
            
          </div>
        </div>
        
      </div>
    </div>
  );
};
export default Sidebar;

const AuthShowcase: React.FC = () => {
  const { data: sessionData } = useSession();

  return (
    <div className="flex flex-col items-center justify-center gap-4">
      <p className="text-indigo-50 my-0">
        {sessionData && <span>{sessionData?.user?.name}</span>}
        
      </p>
      <Link href="/" className="text-indigo-200 text-sm">
              View Profile
            </Link>
      <button
        className="mt-4 inline-block rounded border border-white px-4 py-2 text-sm leading-none text-white hover:border-transparent hover:bg-white hover:text-teal-500 lg:mt-0"
        onClick={sessionData ? () => signOut() : () => signIn()}
      >
        {sessionData ? "Sign out" : "Login"}
      </button>
    </div>
  );
};