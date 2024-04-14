import React from "react";
import {
  CalendarIcon,
  FolderIcon,
  HomeIcon,
  UserGroupIcon,
  GlobeAmericasIcon,
  WrenchScrewdriverIcon,
  ArrowTopRightOnSquareIcon,
  ArrowDownOnSquareIcon,
  EnvelopeOpenIcon,
  ArrowTrendingUpIcon,
  RocketLaunchIcon,
  CalendarDaysIcon,
  ChatBubbleLeftEllipsisIcon,
} from "@heroicons/react/24/outline";
import { NavItem } from "./Sidebar";

export const defaultNavItems: NavItem[] = [
  {
    label: "Dashboard",
    href: "/",
    icon: <HomeIcon className="w-6 h-6" />,
  
  },
  {
    label: "Inventory",
    href: "/inventory",
    icon: <FolderIcon className="w-6 h-6" />,

  },
  {
    label: "Calendar",
    href: "/calendar",
    icon: <CalendarDaysIcon className="w-6 h-6" />,

  },
  {
    label: "Mail",
    href: "/mail",
    icon: <EnvelopeOpenIcon className="w-6 h-6" />,

  },
  {
    label: "Marketing",
    href: "/marketing",
    icon: <CalendarIcon className="w-6 h-6" />,

  },
  {
    label: "Social",
    href: "/social",
    icon: <RocketLaunchIcon className="w-6 h-6" />,

  },
  {
    label: "Analytics",
    href: "/analytics",
    icon: <ArrowTrendingUpIcon className="w-6 h-6" />,

  },
  {
    label: "Import",
    href: "/import",
    icon: <ArrowDownOnSquareIcon className="w-6 h-6" />,

  },
  {
    label: "Export",
    href: "/export",
    icon: <ArrowTopRightOnSquareIcon className="w-6 h-6" />,

  },
  {
    label: "Settings",
    href: "/settings",
    icon: <WrenchScrewdriverIcon className="w-6 h-6" />,

  },
  {
    label: "Research",
    href: "/research",
    icon: <GlobeAmericasIcon className="w-6 h-6" />,

  },
  {
    label: "Team",
    href: "/team",
    icon: <UserGroupIcon className="w-6 h-6" />,

  },
  {
    label: "Help",
    href: "/chatbot",
    icon: <ChatBubbleLeftEllipsisIcon className="w-6 h-6" />,

  },
];