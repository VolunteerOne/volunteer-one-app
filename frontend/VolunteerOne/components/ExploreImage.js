import React from "react";
import {
  StyleSheet,
  Text,
  View,
  Dimensions,
  ImageBackground,
  Animated,
  PanResponder,
} from "react-native";

const ExploreImage = ({ item }) => {
    return (
    <ImageBackground
        style={{
        flex: 1,
        height: null,
        width: null,
        resizeMode: "cover",
        borderRadius: 20,
        }}
        source={item.uri} >
        <View style = {styles.overlay}>
        <Text style = {styles.titleText}>Help pack disaster relief bags</Text>
        <Text style = {styles.eventDetailsText}>{item.author} </Text>


        </View>
    </ImageBackground>
    );
  };
  
  const styles = StyleSheet.create({
    overlay: {
        marginTop: 475,
        backgroundColor: "#FFFFFF",
        height: 175,
        opacity: 0.8,
        borderRadius: 5
    },
    titleText: {
        fontSize: 20,
        fontWeight: "bold",
        color: "#32325D",
        marginTop: 10,
        marginLeft: 15
      },
      eventDetailsText: {
        fontSize: 14,
        color: "#32325D",
        marginTop: 10,
        marginLeft: 15
      },
    shadowProp: {
      shadowColor: "#171717",
      shadowOffset: { width: -2, height: 4 },
      shadowOpacity: 0.2,
      shadowRadius: 3,
    },
  });
  
  export default ExploreImage;
  


