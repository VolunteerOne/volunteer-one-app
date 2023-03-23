import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';

import { PostWithoutImage, PostWithImage, PostImageCard } from '../../components';
import posts from '../../constants/posts';

const { width } = Dimensions.get('screen');

class Home extends React.Component {
  renderPosts = () => {
    var postsList = posts.map(function (data) {
      return <PostImageCard key={data["id"]} data={data} />;
    });

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          {postsList}
        </Block>
      </ScrollView>
    );
  };

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderPosts()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  posts: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default Home;
