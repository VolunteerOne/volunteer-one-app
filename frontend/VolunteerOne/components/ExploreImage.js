import React from "react";
import {
  StyleSheet,
  Text,
  View,
  Dimensions,
  Image,
  Animated,
  PanResponder,
} from "react-native";

const ExploreImage = ({ item }) => {
    return (
    <Image
        style={{
        flex: 1,
        height: null,
        width: null,
        resizeMode: "cover",
        borderRadius: 20,
        }}
        source={item.uri}
    />
    );
  };
  
  const styles = StyleSheet.create({
    card: {
      backgroundColor: "#FFFFFF",
      width: "100%",
      minWidth: "100%",
      borderRadius: 10,
      padding: 15,
      margin: 10,
    },
    shadowProp: {
      shadowColor: "#171717",
      shadowOffset: { width: -2, height: 4 },
      shadowOpacity: 0.2,
      shadowRadius: 3,
    },
  });
  
  export default ExploreImage;
  


