import { MaterialCommunityIcons } from "@expo/vector-icons";
import React from "react";
import Svg, { Path } from "react-native-svg";

export default function SlackIcon() {
  return (
    <Svg width={15} height={15} viewBox="0 0 15 15">
      <Path
        d="M10.5 7.5V6A1.5 1.5 0 0112 7.5h-1.5zm0 0h-3m3 0V2a1.5 1.5 0 00-3 0v5.5m0 0v-3m0 3H2a1.5 1.5 0 010-3h5.5m0 3H13a1.5 1.5 0 000 3H7.5m0-3v3m0-3h-3m3 0V13a1.5 1.5 0 01-3 0V7.5m3-3V3A1.5 1.5 0 106 4.5h1.5zm0 6H9A1.5 1.5 0 017.5 12v-1.5zm-3-3V9A1.5 1.5 0 113 7.5h1.5z"
        fill="black"
      />
    </Svg>
  );
}
